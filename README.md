# github.com/smarty/klaxon/v2

Alerting utilities for Go programs.

## What's a 'klaxon'?

By definition it's a device used to produce an audio signal for alerting purposes (i.e. a car horn).

- https://en.wikipedia.org/wiki/Vehicle_horn#Klaxon

In the [Star Trek](https://en.wikipedia.org/wiki/Star_Trek) universe it is the oft-repeated ["red-alert" sound](https://memory-alpha.fandom.com/wiki/Klaxon), accompanied by flashing red lights all over the ship or station experiencing a state of heightened alert.

## How does this package work?

This package defines a small number of severity levels, which correspond with distinct counter metrics (implementation not provided, but you could use [github.com/smarty/metrics/v2](https://pkg.go.dev/github.com/smartystreets/metrics/v2)).

This package also defines a few alert escalation strategies to calculate a severity based on a series of events (timestamps).

Finally, this package provides a 'sensor' abstraction to tie all of the above together. Provide one of those to your application's monitoring component and you have a convenient signaling mechanism, offering deterministic escalation of alert severity communicated to your infrastructure.

## SMARTY DISCLAIMER

Subject to the terms of the associated license agreement, this software is freely available for your use. This software is FREE, AS IN PUPPIES, and is a gift. Enjoy your new responsibility. This means that while we may consider enhancement requests, we may or may not choose to entertain requests at our sole and absolute discretion.