package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint64 = 4294967295
    v2 := []uint32{2147483648, 764}
    var v3 bool = (255 <= 32767)
    var v4 int64 = -1
    var v5 int64 = v4
    v6 := GenS1{F0: 989, F1: (v4 >> ((1) & 63))}
    v5 = (v4 % ((9223372036854775807) | 1))
    var v7 int64 = (v4 & v4)
    if ((v3 || (3.14159 < (0.0 / 1.0))) && v3) {
        var v8 uint64 = 2
        _ = v8
    } else {
        switch -9223372036854775808 {
        case -9223372036854775808:
            _ = len(v2)
        case -604:
            v1 = (v1 * v1)
        default:
            for fi9 := 0; fi9 < 13; fi9++ {
                if len(v2) > 0 {
                    v2[0] = 0
                }
                v7 = -225
                _ = fi9
            }
        }
        v10 := GenS2{F0: 0}
        v4 -= (v7 << ((4294967295) & 63))
        _ = v10
    }
    var v11 int32 = 1
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 int16 = -86
            switch ("" + "a") {
            case "\n\t":
                var v15 int16 = -1
                _ = v15
            case "hello":
                fmt.Println(v7)
            case "foo bar":
                if v3 {
                    if v3 {
                        var v16 uint32 = 1
                        fmt.Println(v11, v16, v1)
                        _ = v16
                    }
                    var v17 uint8 = 2
                    _ = len(v2)
                    _ = v17
                }
            }
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v5 = v4
    var v18 int16 = 620
    switch 0 {
    case -513:
        if true {
            var v19 float64 = 1e10
            var v20 uint8 = 1
            fmt.Println(v5, v7, v18)
            _ = v19
            _ = v20
        } else {
            var v21 float64 = 1.5
            var v22 string = ("foo bar" + (("\n\t" + ("hello" + ("hello" + ((((("foo bar" + "") + (("hello" + "a") + "foo bar")) + "") + "foo bar") + (("foo bar" + "") + "hello"))))) + ("foo bar" + ("" + ("a" + "a")))))
            switch 1 {
            case -768:
                var v23 uint64 = (v1 >> ((1) & 63))
                _ = v23
            case 128:
                v24 := GenS1{F0: v11, F1: (v4 + -913)}
                _ = v24
            case 388:
                for fi25 := 0; fi25 < 20; fi25++ {
                    v3 = (0 == -1)
                    switch fi25 {
                    case -299:
                        var v26 uint64 = 532
                        _ = v26
                    }
                    var v27 uint16 = 871
                    _ = fi25
                    _ = v27
                }
            }
            _ = v21
            _ = v22
        }
    default:
        for fi28 := 0; fi28 < 20; fi28++ {
            var v29 float64 = 0.0
            switch 128 {
            case -9223372036854775808:
                v30 := GenS2{F0: (v18 | v18)}
                _ = v30
            }
            v18 = v18
            _ = fi28
            _ = v29
        }
    }
    v31 := GenS2{F0: -328}
    switch ("hello" + "a") {
    case "foo bar":
        var v32 string = ""
        _ = v32
    case "":
        var v33 string = ""
        _ = v33
    default:
        for fi34 := 0; fi34 < 11; fi34++ {
            switch (fi34 & fi34) {
            case 9223372036854775807:
                if (4294967295 == 0) {
                    var v35 int = 336
                    _ = v35
                }
            case 127:
                var v36 float64 = float64(-2.25)
                _ = v36
            }
            _ = fi34
        }
    }
    switch -9223372036854775808 {
    case -683:
        v37 := GenS1{F0: (^(v11)), F1: (v7 * (v7 << ((53) & 63)))}
        _ = v37
    case -2147483648:
        if (v18 >= (-(v18))) {
            var v38 int32 = (^(v11))
            if false {
                v39 := GenS1{F0: v11, F1: (^(v4))}
                _ = v39
            } else {
                var v40 string = "\n\t"
                switch "hello" {
                case "hello":
                    v4 += 573
                case "foo bar":
                    var v41 float32 = (3.14159 / 0.1)
                    _ = v41
                }
                var v42 int8 = 100
                _ = v40
                _ = v42
            }
            _ = v38
        }
    }
    var v43 bool = true
    v44 := GenS1{F0: v11, F1: (v5 >> ((486) & 63))}
    fmt.Println(v5, v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
