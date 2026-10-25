package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 uint16; F2 int16 }

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]string)
    var v3 float32 = 0.1
    v4 := GenS1{F0: 32}
    switch "foo bar" {
    case "hello":
        v5 := GenS2{F0: (v1 & 1)}
        _ = v5
    case "\n\t":
        fmt.Println(v1, v3)
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int64 = -1
        v8 := []int{fi6}
        fmt.Println(v1, v7, v7)
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: (v1 + (v1 % ((0) | 1)))}
    for fi10 := 0; fi10 < 11; fi10++ {
        if ((!(!true)) && false) {
            for fi11 := 0; fi11 < 15; fi11++ {
                switch ("a" + "") {
                case "hello":
                    v12 := []int{(fi11 ^ -624), -1}
                    _ = v12
                case "foo bar":
                    v13 := GenS3{F0: ((1.5 + ((0.1 - float64(v3)) + -2.25)) == 1.5)}
                    _ = v13
                default:
                    var v14 uint32 = 536
                    _ = v14
                }
                _ = fi11
            }
        }
        v15 := GenS3{F0: true}
        switch (fi10 - (fi10 >> ((741) & 63))) {
        case -2147483648:
            var v16 float64 = float64(float32((1e10 + 1e10)))
            _ = v16
        case 4294967295:
            v1 += (v1 << ((65535) & 63))
        case 732:
            if false {
                v17 := GenS3{F0: (65535 < 2147483648)}
                switch (("" + ("\n\t" + "\n\t")) + "") {
                case "\n\t":
                    var v18 float32 = 1.5
                    _ = v18
                case "hello":
                    v19 := GenS4{F0: (-(v1)), F1: 65535, F2: v1}
                    _ = v19
                }
                _ = v17
            }
        default:
            for fi20 := 0; fi20 < 18; fi20++ {
                if false {
                    v3 -= v3
                    fmt.Println(fi20, fi10)
                    v1 = (^(v1))
                } else {
                    var v21 int16 = 0
                    _ = v21
                }
                v22 := []int{}
                v2[2147483648] = "foo bar"
                _ = fi20
                _ = v22
            }
        }
        _ = fi10
        _ = v15
    }
    var v23 uint32 = 298
    switch 32768 {
    case 4294967295:
        var v24 float64 = (0.1 / 1.0)
        _ = v24
    case 128:
        v23 = v23
    default:
        for fi25 := 0; fi25 < 5; fi25++ {
            var v26 uint64 = 1
            for fi27 := 0; fi27 < 3; fi27++ {
                var v28 int = fi25
                _ = fi27
                _ = v28
            }
            fmt.Println(v1)
            _ = fi25
            _ = v26
        }
    }
    var v29 uint8 = 42
    v30 := GenS3{F0: (false || (32767 != 141))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v9
}
