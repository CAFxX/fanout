package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
func genFn0(p0 uint8, p1 uint8, p2 int64) int32 {
    v7 := GenS1{F0: false, F1: 42}
    for fi8 := 0; fi8 < 9; fi8++ {
        if (fi8 != 127) {
            var v9 float32 = (3.14159 * 3.14159)
            _ = v9
        } else {
            v10 := []int{}
            p2 += (p2 + p2)
            var v11 int8 = -48
            _ = v10
            _ = v11
        }
        if true {
            for fi12 := 0; fi12 < 1; fi12++ {
                if ((0.1 - float64(1e10)) > (1.5 + 3.14159)) {
                    var v13 uint16 = 487
                    var v14 int16 = int16(fi8)
                    _ = v13
                    _ = v14
                } else {
                    p1 = (^(p1))
                }
                var v15 bool = (false || (4294967295 > 2))
                _ = fi12
                _ = v15
            }
        }
        v16 := GenS1{F0: true, F1: int16(p1)}
        _ = fi8
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (42 <= 799)
    v3 := []uint32{42, 0, 736}
    var v4 uint32 = 2
    var v5 string = "\n\t"
    v6 := []string{}
    var v17 uint32 = v4
    v4 -= v17
    if v2 {
        var v18 int8 = 1
        if ((-9223372036854775808 > 0) || (v2 || false)) {
            if v2 {
                for fi19 := 0; fi19 < 18; fi19++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi20 := 0; fi20 < 11; fi20++ {
                        var v21 int32 = 2147483647
                        v4 -= 195
                        _ = fi20
                        _ = v21
                    }
                    v4 = 2147483647
                    _ = fi19
                }
                var v22 int8 = v18
                var v23 string = ""
                _ = v22
                _ = v23
            }
            switch v5 {
            case "a":
                fmt.Println(v2, v17, v18)
            default:
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 string = v5
                    _ = fi24
                    _ = v25
                }
            }
        }
        _ = v18
    }
    v26 := GenS1{F0: (128 != 35), F1: -1}
    for fi27 := 0; fi27 < 5; fi27++ {
        for fi28 := 0; fi28 < 10; fi28++ {
            fmt.Println(v5, v17, fi28)
            _ = len(v6)
            v29 := GenS1{F0: (false || v2), F1: 58}
            _ = fi28
            _ = v29
        }
        fmt.Println(v17)
        _ = fi27
    }
    v4 = v4
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v30 := GenS1{F0: true, F1: 0}
    switch -9223372036854775808 {
    case 128:
        var v31 uint32 = 1
        _ = v31
    case 32768:
        switch -9223372036854775808 {
        case 905:
            fmt.Println(v5, v4)
        }
    }
    if false {
        v32 := v1[v5]
        _ = len(v3)
        _ = v32
    } else {
        var v33 float64 = (((-2.25 / 1.5) - 0.1) * 1e10)
        switch ("hello" + "") {
        case "hello":
            var v34 int32 = 260
            _ = v34
        default:
            switch v5 {
            case "a":
                switch 127 {
                case -33:
                    if (v2 && (65535 <= 42)) {
                        delete(v1, v5)
                    } else {
                        v4 += (v17 >> ((1) & 63))
                    }
                case 0:
                    v5 = v5
                }
            case "foo bar":
                fmt.Println(v17, v33)
            case "":
                v1[((v5 + "") + "hello")] = v4
            }
        }
        _ = v33
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v35 := v1[("a" + "hello")]
    var v36 string = v5
    for fi37 := 0; fi37 < 8; fi37++ {
        fmt.Println(v35, v2)
        v5 = v36
        _ = fi37
    }
    var v38 int8 = 42
    switch (v36 + ("a" + (("foo bar" + v36) + v5))) {
    case "hello":
        var v39 int8 = (^(v38))
        _ = v39
    case "\n\t":
        var v40 uint8 = 0
        _ = v40
    default:
        v41 := GenS1{F0: v2, F1: 0}
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
