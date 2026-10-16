package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int8 }

func main() {
    var v1 bool = (!true)
    v2 := make(map[uint32]string)
    var v3 uint64 = 0
    v4 := GenS1{F0: 2, F1: -442, F2: 34}
    var v5 int32 = 1
    switch 96 {
    case 127:
        var v6 int32 = -242
        _ = v6
    default:
        var v7 int = 127
        _ = v7
    }
    v3 -= v3
    _ = len(v2)
    for fi8 := 0; fi8 < 9; fi8++ {
        if (42 >= 127) {
            switch ("foo bar" + "foo bar") {
            case "\n\t":
                v5 += 765
            case "":
                v9 := []int{}
                _ = v9
            }
            var v10 int8 = -96
            _ = v10
        }
        if (355 >= 2) {
            v5 = (v5 * 2147483647)
            for fi11 := 0; fi11 < 7; fi11++ {
                fmt.Println(v3)
                _ = fi11
            }
            v1 = (520 == 255)
        } else {
            for fi12 := 0; fi12 < 3; fi12++ {
                switch int(v5) {
                case 128:
                    v1 = ((!((42 < -669) || true)) && v1)
                case 32767:
                    v3 = v3
                default:
                    v13 := []int{int(v3), (fi8 << ((149) & 63)), fi8}
                    _ = v13
                }
                v14 := GenS1{F0: 4294967295, F1: 732, F2: 1}
                if (v1 && (!v1)) {
                    v15 := GenS1{F0: 1, F1: 42, F2: -1}
                    v16 := []int{}
                    v5 = -1
                    _ = v15
                    _ = v16
                } else {
                    v17 := []int{32768, fi8, (fi8 >> ((65535) & 63)), fi8}
                    v5 += v5
                    v1 = v1
                    _ = v17
                }
                _ = fi12
                _ = v14
            }
        }
        _ = fi8
    }
    _ = len(v2)
    v1 = (!(!(v1 && (true && v1))))
    fmt.Println(v3)
    v18 := []int{0}
    fmt.Println(v1, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
