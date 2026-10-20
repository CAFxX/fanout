package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int32; F1 int8 }
func genFn0(p0 int64, p1 int32, p2 uint) int16 {
    var v4 bool = (((uint32(p2) == 2147483647) || true) && (42 > 65535))
    var v5 bool = true
    v6 := GenS1{F0: p2}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 float64 = 0.0
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 int8 = 12
            _ = fi9
            _ = v10
        }
        v11 := []int{(fi7 ^ fi7), int(p1), fi7}
        _ = fi7
        _ = v11
        _ = v8
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 397
}

func main() {
    var v1 int16 = -143
    var v2 int64 = 42
    v3 := []string{"\n\t", "foo bar", ("foo bar" + "foo bar"), (("a" + ("foo bar" + ("a" + "a"))) + "")}
    v12 := GenS1{F0: 65535}
    var v13 int16 = (v1 + genFn0(v2, 0, 255))
    if false {
        v1 = v1
        v14 := GenS1{F0: 255}
        var v15 string = ("foo bar" + "")
        _ = v14
        _ = v15
    }
    var v16 uint16 = 294
    if true {
        v2 = int64(v16)
        var v17 float32 = 1e10
        var v18 int64 = -1
        _ = v17
        _ = v18
    } else {
        v13 += 1
        var v19 int8 = 36
        switch int(v1) {
        case 128:
            for fi20 := 0; fi20 < 11; fi20++ {
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                switch (-(fi20)) {
                case 2147483647:
                    var v21 uint32 = 42
                    _ = v21
                case -62:
                    switch fi20 {
                    case 127:
                        v22 := GenS1{F0: 701}
                        _ = v22
                    default:
                        var v23 int = -769
                        _ = v23
                    }
                }
                _ = fi20
            }
        }
        _ = v19
    }
    fmt.Println(v2)
    if (!(!(4294967295 <= 77))) {
        if len(v3) > 0 {
            v3[0] = ("\n\t" + "")
        }
        v24 := GenS2{F0: true, F1: 2147483647, F2: 42}
        _ = v24
    } else {
        v25 := GenS1{F0: 815}
        if true {
            if false {
                v13 = 0
                if false {
                    var v26 int32 = -290
                    v1 -= (-(v13))
                    _ = v26
                }
            }
            switch -284 {
            case 0:
                v27 := GenS2{F0: (42 == 0), F1: 2, F2: 2}
                _ = v27
            case -1:
                for fi28 := 0; fi28 < 15; fi28++ {
                    v2 = v2
                    v1 += -1
                    _ = fi28
                }
            case 9223372036854775807:
                if (14 < 1) {
                    v1 -= 835
                    for fi29 := 0; fi29 < 14; fi29++ {
                        if true {
                            var v30 float64 = 3.14159
                            _ = v30
                        }
                        fmt.Println(v2, v16)
                        var v31 uint = 1
                        _ = fi29
                        _ = v31
                    }
                } else {
                    for fi32 := 0; fi32 < 12; fi32++ {
                        var v33 float64 = ((0.0 + 1e10) + 3.14159)
                        _ = fi32
                        _ = v33
                    }
                }
            default:
                var v34 bool = true
                _ = v34
            }
        }
        v16 = v16
        _ = v25
    }
    var v35 int = -2147483648
    v35 -= (v35 | -2147483648)
    v3 = append(v3, "a")
    if false {
        var v36 int8 = 0
        v3 = append(v3, "\n\t")
        v36 -= (v36 - v36)
        _ = v36
    }
    _ = len(v3)
    v13 = v13
    v35 = v35
    if false {
        for fi37 := 0; fi37 < 20; fi37++ {
            switch (v35 << ((716) & 63)) {
            case 42:
                for fi38 := 0; fi38 < 20; fi38++ {
                    fmt.Println(v35)
                    if ((fi37 / (((fi37 << ((uint(v16)) & 63))) | 1)) == fi38) {
                        var v39 uint = 18446744073709551615
                        v40 := GenS3{F0: -1, F1: 1}
                        _ = v39
                        _ = v40
                    } else {
                        var v41 int64 = v2
                        _ = v41
                    }
                    v1 -= v13
                    _ = fi38
                }
            case -2147483648:
                var v42 int = 42
                _ = v42
            case 127:
                fmt.Println(v13, v16, v1)
            }
            v2 += 0
            _ = fi37
        }
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v35
}
