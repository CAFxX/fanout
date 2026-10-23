package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 int = 81
    var v3 uint32 = 724
    var v4 int16 = -32768
    v4 = (v4 >> ((42) & 63))
    v5 := GenS1{F0: v2, F1: 42}
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(fi6, fi6, v4)
        _ = fi6
    }
    if (!false) {
        v7 := GenS2{F0: true}
        v3 -= v3
        var v8 int16 = 1000
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        v3 += v3
        var v10 bool = false
        if (!false) {
            var v11 float32 = 0.1
            var v12 int = (fi9 << ((255) & 63))
            v13 := []int{(^(fi9))}
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    var v14 string = ""
    delete(v1, (v2 << ((4294967295) & 63)))
    v14 = "hello"
    var v15 int = -9223372036854775808
    v4 = v4
    if true {
        if (((v14 + v14) < (((v14 + v14) + v14) + v14)) || true) {
            switch (-(v2)) {
            case -104:
                switch v14 {
                case "a":
                    for fi16 := 0; fi16 < 7; fi16++ {
                        if true {
                            var v17 uint32 = (v3 ^ v3)
                            var v18 int8 = 42
                            _ = v17
                            _ = v18
                        } else {
                            var v19 int = v2
                            var v20 uint16 = 676
                            _ = v19
                            _ = v20
                        }
                        _ = fi16
                    }
                case "foo bar":
                    if (!((42 >= 1) && (((1.5 * float64((0.1 + 3.14159))) == (0.1 * 3.14159)) || (false || (!true))))) {
                        for fi21 := 0; fi21 < 5; fi21++ {
                            v14 = (v14 + ((v14 + v14) + v14))
                            _ = fi21
                        }
                    }
                }
            case 42:
                v2 += (v15 + v2)
            default:
                var v22 int32 = 418
                _ = v22
            }
        } else {
            v3 = 42
            var v23 int32 = -931
            _ = v23
        }
        if false {
            v2 += -62
            var v24 uint32 = v3
            var v25 uint64 = 926
            _ = v24
            _ = v25
        }
        v2 -= 0
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        switch (v2 << ((1) & 63)) {
        case -1:
            var v27 int64 = 43
            _ = v27
        case 530:
            v28 := GenS3{F0: 127, F1: 9223372036854775807}
            _ = v28
        case -676:
            if (!(!true)) {
                v29 := []int{fi26, -587, (v15 ^ fi26)}
                _ = v29
            }
        default:
            v30 := GenS2{F0: (18446744073709551615 == 946)}
            _ = v30
        }
        _ = fi26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
