package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []uint32{903, 1}
    var v2 int32 = 782
    var v3 int16 = 783
    var v4 bool = false
    v5 := []string{"hello"}
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    v2 = (v2 * (v2 >> ((42) & 63)))
    var v6 int = 127
    v4 = (-632 >= -995)
    v6 = v6
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int32 = 0
        v8 = 0
        fmt.Println(v3, v4, v2)
        _ = fi7
        _ = v8
    }
    v4 = v4
    var v9 int32 = v2
    var v10 uint16 = 759
    if (42 == 306) {
        switch "hello" {
        case "a":
            v6 += v6
        case "foo bar":
            switch (v6 * 128) {
            case -262:
                v2 = (v2 / ((190) | 1))
            case 619:
                switch "a" {
                case "\n\t":
                    fmt.Println(v3, v10, v9)
                }
            case 128:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        case "\n\t":
            for fi11 := 0; fi11 < 6; fi11++ {
                for fi12 := 0; fi12 < 19; fi12++ {
                    v13 := GenS1{F0: int8(v10)}
                    v5 = append(v5, "a")
                    var v14 int64 = 1
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                var v15 float64 = (((3.14159 + (0.0 * -2.25)) + 3.14159) * (float64(1e10) - 1e10))
                if v4 {
                    var v16 float64 = (0.1 * 0.0)
                    v16 = 0.1
                    _ = v16
                }
                _ = fi11
                _ = v15
            }
        }
        var v17 int16 = v3
        v18 := GenS1{F0: -128}
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 uint16 = (v10 + (^(v10)))
            var v21 int64 = 1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 int16 = -32768
        v9 -= v9
        _ = v22
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        var v24 int32 = -343
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 float32 = (-2.25 * 0.1)
            v24 += (v24 << ((65535) & 63))
            _ = fi25
            _ = v26
        }
        v5 = append(v5, "hello")
        _ = fi23
        _ = v24
    }
    var v27 int32 = 0
    if (v4 || v4) {
        v5 = append(v5, ("\n\t" + ("foo bar" + (((("" + "\n\t") + "\n\t") + "a") + ("hello" + "foo bar")))))
    } else {
        if (3.14159 < ((0.1 / 3.14159) / 1.0)) {
            var v28 float32 = (0.0 / -2.25)
            _ = v28
        }
    }
    v2 = -119
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
