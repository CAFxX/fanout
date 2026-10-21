package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint8 = 65
    var v2 uint8 = v1
    v3 := make(map[uint32]string)
    v2 += v2
    for fi4 := 0; fi4 < 3; fi4++ {
        switch "hello" {
        case "foo bar":
            switch (-(fi4)) {
            case -2147483648:
                if true {
                    fmt.Println(v1, v1, fi4)
                    v5 := GenS1{F0: "foo bar", F1: 1000, F2: (^(fi4))}
                    v6 := GenS1{F0: "\n\t", F1: 42, F2: fi4}
                    _ = v5
                    _ = v6
                }
            case 9223372036854775807:
                if (216 <= 65535) {
                    var v7 int8 = 127
                    for fi8 := 0; fi8 < 3; fi8++ {
                        v9 := []int{}
                        v1 -= 127
                        _ = fi8
                        _ = v9
                    }
                    v1 = v2
                    _ = v7
                }
            case 128:
                v1 += (v2 % (((v1 >> ((42) & 63))) | 1))
            default:
                for fi10 := 0; fi10 < 20; fi10++ {
                    v11 := GenS1{F0: "\n\t", F1: -647, F2: 0}
                    v12 := []int{0, (fi4 ^ (fi4 << ((0) & 63)))}
                    v1 = (v2 & (v2 - v1))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
        }
        _ = fi4
    }
    v2 -= v1
    v13 := GenS1{F0: (((("\n\t" + "hello") + "a") + (("\n\t" + (("a" + (("" + "hello") + (("" + "a") + "\n\t"))) + "\n\t")) + "hello")) + ("\n\t" + ("" + "foo bar"))), F1: 32767, F2: -877}
    v14 := GenS2{F0: -1}
    v2 -= (v2 | (v2 + (v1 / ((v1) | 1))))
    var v15 int8 = 1
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 string = "hello"
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
