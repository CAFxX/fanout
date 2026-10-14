package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 string }

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 42
    var v3 int64 = 0
    if (true && true) {
        switch ("hello" + ("\n\t" + (("foo bar" + "a") + "foo bar"))) {
        case "foo bar":
            var v4 uint8 = 34
            _ = v4
        case "":
            v2 = v2
        case "a":
            if (255 > 2) {
                var v5 uint16 = 418
                v6 := []int{69}
                _ = v5
                _ = v6
            }
        default:
            for fi7 := 0; fi7 < 1; fi7++ {
                if false {
                    v8 := []int{}
                    _ = v8
                }
                _ = fi7
            }
        }
        v1 = (v1 * v1)
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS1{F0: 1}
        _ = fi9
        _ = v10
    }
    var v11 bool = false
    v1 += (v1 + v1)
    for fi12 := 0; fi12 < 4; fi12++ {
        switch "hello" {
        case "a":
            v11 = false
        case "hello":
            v11 = (fi12 < 9223372036854775807)
        case "\n\t":
            switch 32767 {
            case -546:
                switch fi12 {
                case 32768:
                    if (v1 >= 42) {
                        var v13 int32 = 244
                        _ = v13
                    }
                case 74:
                    v2 -= (v2 - (v2 >> ((42) & 63)))
                case 128:
                    v14 := []int{fi12, fi12, fi12}
                    _ = v14
                }
            }
        }
        if ((v3 >> ((1) & 63)) > 0) {
            v15 := GenS2{F0: (v1 << ((65535) & 63)), F1: ("hello" + "")}
            var v16 int8 = 127
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 int = (fi17 & fi17)
            _ = fi17
            _ = v18
        }
        _ = fi12
    }
    v11 = v11
    v19 := []int{1}
    switch "a" {
    case "":
        switch 9223372036854775807 {
        case 127:
            if true {
                var v20 int8 = 127
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                _ = v20
            } else {
                var v21 uint8 = 175
                v21 -= uint8(v2)
                _ = v21
            }
        default:
            if len(v19) > 0 {
                v19[0] = 812
            }
        }
    default:
        if len(v19) > 0 {
            v19[0] = -575
        }
    }
    var v22 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v22
    _ = v3
}
