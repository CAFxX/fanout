package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 uint32; F2 string }
func genFn0(p0 uint16, p1 int16, p2 uint64) uint {
    v6 := GenS1{F0: 42}
    for fi7 := 0; fi7 < 7; fi7++ {
        if ((p1 + p1) == p1) {
            p0 = 32767
            var v8 uint = 1
            var v9 uint32 = 4294967295
            _ = v8
            _ = v9
        } else {
            v10 := GenS2{F0: p2}
            p2 -= 4294967295
            _ = v10
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}
func genFn1(p0 string, p1 int64) bool {
    if true {
        p0 = "a"
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint16 = 32767
            switch (^(fi11)) {
            case 127:
                var v13 uint = 301
                _ = v13
            default:
                if (!false) {
                    if false {
                        for fi14 := 0; fi14 < 16; fi14++ {
                            var v15 int64 = p1
                            v16 := GenS3{F0: fi14, F1: (v12 <= v12)}
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                    switch fi11 {
                    case 2147483647:
                        var v17 uint64 = 511
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 10; fi18++ {
                        var v19 string = "foo bar"
                        _ = fi18
                        _ = v19
                    }
                }
            }
            _ = fi11
            _ = v12
        }
    } else {
        var v20 string = p0
        v21 := GenS3{F0: 32767, F1: (p0 <= "a")}
        _ = v20
        _ = v21
    }
    v22 := []int{641, 813, int(p1)}
    var v23 bool = (!(65535 != 2))
    p0 = (p0 + (p0 + "foo bar"))
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return ((-1 != 42) && v23)
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 2147483647
    var v3 string = (("a" + "foo bar") + ("a" + "\n\t"))
    v4 := make(map[string]uint32)
    var v5 uint32 = 65535
    v24 := []int{}
    if (2 >= 855) {
        v3 = v3
        var v25 int16 = 935
        _ = v25
    }
    switch "" {
    case "foo bar":
        fmt.Println(v3, v3)
    }
    v26 := GenS1{F0: 113}
    v2 -= v2
    for fi27 := 0; fi27 < 4; fi27++ {
        switch v3 {
        case "foo bar":
            var v28 float64 = 3.14159
            _ = v28
        default:
            v3 = v3
        }
        var v29 int64 = 947
        _ = fi27
        _ = v29
    }
    v30 := GenS4{F0: 694}
    switch ("" + (v3 + v3)) {
    case "foo bar":
        switch 250 {
        case -632:
            v2 = v2
        }
    default:
        switch v2 {
        case 480:
            v31 := GenS3{F0: v2, F1: (746 != 4294967295)}
            _ = v31
        case -377:
            var v32 float32 = 0.0
            _ = v32
        case -458:
            v2 = v2
        default:
            var v33 bool = true
            _ = v33
        }
    }
    v2 += (v2 << ((279) & 63))
    for fi34 := 0; fi34 < 17; fi34++ {
        v35 := GenS5{F0: 100, F1: (^(v5)), F2: ""}
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        var v37 uint = 319
        v38 := GenS5{F0: 0, F1: v5, F2: ""}
        _ = fi36
        _ = v37
        _ = v38
    }
    if false {
        for fi39 := 0; fi39 < 17; fi39++ {
            v1[(v3 + "")] = "a"
            fmt.Println(v2)
            var v40 bool = (65535 != (v5 | v5))
            _ = fi39
            _ = v40
        }
        v41 := GenS2{F0: 0}
        _ = v41
    } else {
        var v42 int8 = -125
        v43 := v1["foo bar"]
        for fi44 := 0; fi44 < 20; fi44++ {
            v3 = (v43 + ("" + v43))
            _ = fi44
        }
        _ = v42
        _ = v43
    }
    var v45 int8 = -73
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v5
}
