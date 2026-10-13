package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 0
    var v3 int16 = 42
    v3 = v3
    _ = len(v1)
    var v4 uint16 = 65535
    v5 := GenS1{F0: 1, F1: (v2 * (^(v2)))}
    var v6 int32 = 262
    for fi7 := 0; fi7 < 1; fi7++ {
        if (v4 >= (v4 ^ v4)) {
            var v8 int = 128
            _ = v8
        } else {
            switch ("a" + ("a" + "a")) {
            case "foo bar":
                if true {
                    switch ("a" + "a") {
                    case "a":
                        var v9 uint = 255
                        _ = v9
                    case "hello":
                        for fi10 := 0; fi10 < 17; fi10++ {
                            v11 := GenS1{F0: 1, F1: (v2 ^ v2)}
                            var v12 float32 = -2.25
                            if (!true) {
                                var v13 uint32 = (v2 >> ((2) & 63))
                                v3 = -75
                                _ = v13
                            }
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                    case "\n\t":
                        v2 = (v2 & v2)
                    }
                } else {
                    v3 += 0
                }
            default:
                fmt.Println(fi7, v2)
            }
        }
        _ = len(v1)
        _ = fi7
    }
    switch 1 {
    case -294:
        var v14 uint64 = 683
        _ = v14
    case 128:
        var v15 float32 = 1.5
        _ = v15
    case 32768:
        v16 := []int{127, 133, -750}
        _ = v16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        if false {
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS1{F0: 4294967295, F1: v2}
                var v20 uint8 = 0
                fmt.Println(v6)
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
