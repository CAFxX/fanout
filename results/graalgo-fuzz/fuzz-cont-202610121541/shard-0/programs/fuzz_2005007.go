package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint32; F2 bool }
type GenS3 struct { F0 int16; F1 uint64; F2 uint }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int64; F1 bool }
func genFn0(p0 uint32, p1 uint) int64 {
    var v4 uint = 0
    for fi5 := 0; fi5 < 16; fi5++ {
        fmt.Println(v4, p1, fi5)
        _ = fi5
    }
    var v6 bool = false
    for fi7 := 0; fi7 < 13; fi7++ {
        p1 += (p1 * (v4 << ((42) & 63)))
        var v8 uint16 = 770
        if v6 {
            switch "" {
            case "hello":
                fmt.Println(fi7)
            }
            fmt.Println(v4)
        } else {
            if v6 {
                var v9 int64 = 42
                _ = v9
            }
            v4 = v4
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 548
}
func genFn1(p0 int64, p1 bool) int16 {
    var v10 bool = false
    p0 -= p0
    var v11 uint8 = 29
    if true {
        for fi12 := 0; fi12 < 18; fi12++ {
            p0 -= 0
            _ = fi12
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -1
}

func main() {
    var v1 int16 = -340
    var v2 uint = 462
    var v3 string = ("foo bar" + "\n\t")
    var v13 bool = (true && false)
    v3 = v3
    v14 := []int{-9223372036854775808, 32768}
    var v15 uint64 = 2
    if (-2147483648 >= 2147483647) {
        var v16 float32 = 1e10
        _ = v16
    } else {
        switch 111 {
        case -2147483648:
            v1 += v1
        case 0:
            for fi17 := 0; fi17 < 13; fi17++ {
                for fi18 := 0; fi18 < 9; fi18++ {
                    var v19 int = (fi17 >> ((1) & 63))
                    switch (fi17 | (fi17 + (fi18 ^ (fi17 | v19)))) {
                    case 1:
                        v13 = (v3 != "foo bar")
                    case 0:
                        fmt.Println(v3, v19, v13)
                        var v20 uint8 = 133
                        var v21 uint = v2
                        _ = v20
                        _ = v21
                    default:
                        v22 := GenS1{F0: v2}
                        _ = v22
                    }
                    _ = fi18
                    _ = v19
                }
                v15 += 1
                var v23 float32 = 1.5
                _ = fi17
                _ = v23
            }
        case 32768:
            for fi24 := 0; fi24 < 8; fi24++ {
                var v25 int32 = 886
                _ = fi24
                _ = v25
            }
        }
    }
    if len(v14) > 0 {
        v14[0] = 9223372036854775807
    }
    v15 = uint64(v2)
    for fi26 := 0; fi26 < 9; fi26++ {
        if (0.1 < 3.14159) {
            v14 = append(v14, (fi26 % (((fi26 + (fi26 + fi26))) | 1)))
        }
        v2 -= 42
        _ = fi26
    }
    if v13 {
        var v27 int32 = -149
        if len(v14) > 0 {
            v14[0] = -366
        }
        _ = v27
    } else {
        if v13 {
            var v28 int32 = -1
            _ = v28
        } else {
            if true {
                v29 := GenS2{F0: 0, F1: 1, F2: (-48 == 81)}
                _ = v29
            }
        }
        var v30 bool = (!(2 == 2))
        _ = v30
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        for fi32 := 0; fi32 < 20; fi32++ {
            v33 := GenS3{F0: (v1 << ((v2) & 63)), F1: uint64(v2), F2: 255}
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    if len(v14) > 0 {
        v14[0] = 128
    }
    v34 := GenS4{F0: 0, F1: 127}
    for fi35 := 0; fi35 < 2; fi35++ {
        var v36 uint32 = 807
        if (!v13) {
            v37 := GenS4{F0: uint16(v2), F1: fi35}
            _ = v37
        }
        _ = fi35
        _ = v36
    }
    if false {
        v15 += 0
        if (0.1 >= float32(0.1)) {
            v38 := GenS5{F0: 9223372036854775807, F1: (genFn0(2147483647, 255) != 9223372036854775807)}
            v39 := GenS4{F0: 32767, F1: 32768}
            _ = v38
            _ = v39
        }
        if v13 {
            var v40 string = v3
            _ = len(v14)
            var v41 uint8 = 42
            _ = v40
            _ = v41
        }
    } else {
        var v42 int64 = -349
        var v43 int32 = -2147483648
        v1 += (v1 << (((v2 - v2)) & 63))
        _ = v42
        _ = v43
    }
    _ = len(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
}
