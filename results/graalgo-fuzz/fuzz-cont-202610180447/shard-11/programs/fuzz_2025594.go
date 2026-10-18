package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 4294967295
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    switch (v2 >> ((81) & 63)) {
    case 1:
        v2 = (v2 / ((v2) | 1))
    case 0:
        var v5 uint64 = 4294967295
        _ = v5
    case -188:
        v6 := GenS1{F0: true, F1: 1}
        _ = v6
    default:
        v7 := v3[1]
        _ = v7
    }
    v8 := GenS1{F0: (42 < 42), F1: -277}
    if (1 >= 0) {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 bool = (false || (65535 == 392))
            _ = fi9
            _ = v10
        }
    } else {
        v2 -= v2
        v2 += (^(v2))
    }
    if true {
        for fi11 := 0; fi11 < 13; fi11++ {
            fmt.Println(fi11)
            _ = fi11
        }
        switch "a" {
        case "foo bar":
            switch (v2 / ((v2) | 1)) {
            case -809:
                v2 -= 128
            case 4294967295:
                for fi12 := 0; fi12 < 14; fi12++ {
                    v13 := GenS1{F0: (true && true), F1: v2}
                    _ = fi12
                    _ = v13
                }
            }
        case "":
            v2 = (v2 * -9223372036854775808)
        default:
            var v14 uint32 = 106
            _ = v14
        }
    }
    v2 += v2
    v15 := GenS1{F0: false, F1: (v2 & (v2 * (v2 | -867)))}
    v16 := v4["foo bar"]
    for fi17 := 0; fi17 < 11; fi17++ {
        if (false && true) {
            v2 -= fi17
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := []int{fi17, (fi17 | (fi18 | fi18)), (fi17 % ((1) | 1))}
                _ = len(v3)
                var v20 int = v2
                _ = fi18
                _ = v19
                _ = v20
            }
            if (!true) {
                switch v2 {
                case 42:
                    v2 += (fi17 >> ((4294967295) & 63))
                case -2147483648:
                    for fi21 := 0; fi21 < 20; fi21++ {
                        var v22 bool = (128 == 1)
                        _ = fi21
                        _ = v22
                    }
                }
            } else {
                v16 -= (^(v16))
            }
        } else {
            fmt.Println(fi17, fi17, fi17)
            var v23 int = v2
            _ = v23
        }
        switch 400 {
        case 128:
            v2 = fi17
        case 32767:
            fmt.Println(v2, v2, v16)
        default:
            v16 -= (^(v16))
        }
        v2 = (v2 ^ (v2 / ((fi17) | 1)))
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
