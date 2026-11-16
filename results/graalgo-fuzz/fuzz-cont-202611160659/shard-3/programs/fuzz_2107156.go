package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []string{"a", "foo bar"}
    var v2 bool = false
    var v3 string = ("" + "foo bar")
    var v4 float32 = 1e10
    var v5 uint16 = 447
    v6 := []int{-1, 128, -984}
    for fi7 := 0; fi7 < 8; fi7++ {
        v5 += (^(v5))
        switch ((v3 + v3) + v3) {
        case "a":
            v2 = (float32(0.0) >= float32(-2.25))
        case "":
            if v2 {
                var v8 uint32 = 2147483648
                v5 = v5
                v8 = 1
                _ = v8
            } else {
                fmt.Println(v4, fi7, v2)
                v5 += (v5 * (v5 | (v5 >> ((42) & 63))))
            }
        case "\n\t":
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 uint16 = (^(v5))
                v6 = append(v6, (fi9 + (fi9 << ((720) & 63))))
                v5 = (v10 << ((1) & 63))
                var v11 uint = 590
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi7
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 int16 = -247
        if len(v6) > 0 {
            v6[0] = -462
        }
        fmt.Println(fi12)
        _ = fi12
        _ = v13
    }
    switch 127 {
    case 1:
        for fi14 := 0; fi14 < 7; fi14++ {
            if len(v1) > 0 {
                v1[0] = v3
            }
            _ = fi14
        }
    }
    fmt.Println(v3, v4, v5)
    v4 = v4
    v6 = append(v6, 128)
    var v15 uint16 = (v5 + 65535)
    var v16 int = 32767
    v2 = v2
    v17 := GenS1{F0: 42}
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := GenS2{F0: -373}
        _ = fi18
        _ = v19
    }
    var v20 uint = 2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        v22 := GenS1{F0: 67}
        v23 := GenS2{F0: 1000}
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: 62}
    v20 = v20
    fmt.Println(v3)
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
