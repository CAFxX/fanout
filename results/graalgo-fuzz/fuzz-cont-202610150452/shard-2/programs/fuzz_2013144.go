package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 42
    var v3 int = 32768
    var v4 float32 = float32(-2.25)
    var v5 int = (v3 ^ (v3 << ((231) & 63)))
    fmt.Println(v5)
    switch 32767 {
    case 32767:
        fmt.Println(v3, v3, v2)
    case 2147483647:
        v2 = 220
    case 0:
        for fi6 := 0; fi6 < 3; fi6++ {
            v5 = 32767
            var v7 int16 = -32768
            _ = fi6
            _ = v7
        }
    }
    v3 = (v5 + (v5 | 32767))
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 string = "a"
        var v10 int32 = -1
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        var v11 int = (-(v3))
        var v12 int16 = 42
        var v13 int = v5
        _ = v11
        _ = v12
        _ = v13
    } else {
        var v14 int32 = -162
        if true {
            v15 := []int{(v5 & 94), (v5 - (v5 >> ((13) & 63)))}
            if true {
                v16 := GenS1{F0: 255, F1: 448, F2: 2}
                v4 += v4
                _ = v16
            } else {
                v5 += v5
                var v17 uint8 = v2
                for fi18 := 0; fi18 < 18; fi18++ {
                    v4 = v4
                    _ = fi18
                }
                _ = v17
            }
            _ = v15
        } else {
            if false {
                v19 := []int{v5, (^(v3))}
                fmt.Println(v14)
                v20 := v1[-1]
                _ = v19
                _ = v20
            }
            v2 -= v2
        }
        _ = v14
    }
    v5 -= (-(v3))
    v21 := GenS1{F0: 42, F1: 135, F2: 2}
    v3 = v5
    v4 += v4
    var v22 string = "\n\t"
    var v23 uint64 = 2
    if false {
        var v24 uint8 = v2
        for fi25 := 0; fi25 < 8; fi25++ {
            switch (^(fi25)) {
            case 4294967295:
                for fi26 := 0; fi26 < 17; fi26++ {
                    var v27 float32 = float32(0.1)
                    v28 := GenS1{F0: v2, F1: 118, F2: uint16(v23)}
                    v1[v5] = v22
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            if ((!((-(fi25)) < (v3 * (fi25 ^ v5)))) && (((("" + v22) + "") >= (v22 + (v22 + ((v22 + v22) + v22)))) || (255 == 0))) {
                v3 = -448
            }
            v22 = v22
            _ = fi25
        }
        var v29 int64 = -9223372036854775808
        _ = v24
        _ = v29
    } else {
        var v30 uint = 0
        for fi31 := 0; fi31 < 3; fi31++ {
            for fi32 := 0; fi32 < 2; fi32++ {
                var v33 int = -87
                fmt.Println(fi31, v22)
                var v34 int16 = 42
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi31
        }
        for fi35 := 0; fi35 < 5; fi35++ {
            v4 += float32((-2.25 + (1e10 / 1e10)))
            fmt.Println(v22, v4)
            _ = fi35
        }
        _ = v30
    }
    switch v22 {
    case "a":
        var v36 string = v22
        _ = v36
    }
    v22 = ("a" + v22)
    for fi37 := 0; fi37 < 11; fi37++ {
        switch int(v23) {
        case 128:
            v2 -= 0
        case 32767:
            fmt.Println(v4)
        }
        if (127 >= 42) {
            v23 -= 42
        }
        delete(v1, 1)
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
