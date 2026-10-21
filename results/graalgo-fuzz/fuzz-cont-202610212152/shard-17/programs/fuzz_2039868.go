package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 bool; F2 uint }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 int8; F2 int64 }
type GenS7 struct { F0 int32; F1 string }
func genFn0(p0 int, p1 int64) int32 {
    p1 += -277
    if ((false || true) || false) {
        p0 += (p0 / ((32767) | 1))
        var v4 int64 = (-(p1))
        for fi5 := 0; fi5 < 19; fi5++ {
            fmt.Println(v4, fi5, fi5)
            _ = fi5
        }
        _ = v4
    }
    v6 := []int{(p0 | p0), -2147483648, (p0 << ((1) & 63)), p0}
    var v7 uint16 = 255
    p0 = (p0 << ((0) & 63))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 int8, p1 int8) int {
    var v8 bool = ("hello" >= (("\n\t" + "\n\t") + (("\n\t" + "\n\t") + "a")))
    fmt.Println(p0, v8)
    switch -1 {
    case 9223372036854775807:
        p1 = (p1 >> ((985) & 63))
    case 2147483647:
        v8 = v8
    }
    _ = p0
    _ = p1
    _ = v8
    return 32768
}

func main() {
    v1 := []uint32{}
    var v2 string = "foo bar"
    v3 := make(map[int]string)
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: 1}
        _ = fi9
        _ = v10
    }
    v1 = append(v1, 899)
    switch v2 {
    case "hello":
        switch 42 {
        case -2147483648:
            v2 = v2
        }
    case "foo bar":
        fmt.Println(v2)
    case "\n\t":
        var v11 int32 = -1
        _ = v11
    default:
        var v12 uint64 = 18446744073709551615
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        _ = len(v3)
        _ = fi13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        v2 = "\n\t"
        for fi15 := 0; fi15 < 20; fi15++ {
            switch (v2 + v2) {
            case "a":
                v2 = (((v2 + v2) + (v2 + v2)) + "foo bar")
            default:
                if (!(!false)) {
                    _ = len(v1)
                    fmt.Println(v2, fi15, v2)
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    v2 = (("" + v2) + v2)
                    v2 = ""
                    var v16 bool = false
                    _ = v16
                }
            }
            var v17 int32 = 304
            if true {
                v18 := GenS2{F0: fi15}
                v3[(fi14 / ((fi15) | 1))] = (("a" + v2) + v2)
                if true {
                    var v19 int16 = -47
                    var v20 int64 = 861
                    _ = v19
                    _ = v20
                } else {
                    v21 := GenS2{F0: genFn1(98, 100)}
                    _ = v21
                }
                _ = v18
            }
            _ = fi15
            _ = v17
        }
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 float64 = (3.14159 - (-2.25 + float64(float32(-2.25))))
            _ = fi22
            _ = v23
        }
        _ = fi14
    }
    v2 = ""
    v24 := v3[1]
    if (true && false) {
        if (42 >= 0) {
            v24 = v24
            fmt.Println(v24)
            v3[genFn1(100, 1)] = "\n\t"
        } else {
            var v25 bool = (0 == 18446744073709551615)
            v26 := GenS3{F0: 127, F1: v25, F2: 65535}
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 int32 = -371
                _ = fi27
                _ = v28
            }
            _ = v25
            _ = v26
        }
        var v29 int16 = 1000
        for fi30 := 0; fi30 < 6; fi30++ {
            v3[fi30] = ("hello" + "")
            _ = fi30
        }
        _ = v29
    } else {
        var v31 uint32 = 2
        v24 = v24
        var v32 int32 = genFn0(1, -339)
        _ = v31
        _ = v32
    }
    v33 := GenS4{F0: 2147483647, F1: 32767}
    var v34 float64 = (3.14159 * 0.0)
    var v35 int = 42
    if ((!true) && false) {
        if false {
            var v36 int32 = int32(v35)
            for fi37 := 0; fi37 < 7; fi37++ {
                v35 = 4294967295
                _ = fi37
            }
            switch 32767 {
            case 2147483647:
                var v38 float32 = 3.14159
                _ = v38
            }
            _ = v36
        }
        switch (v35 << ((1) & 63)) {
        case 9223372036854775807:
            fmt.Println(v2, v34)
        }
    } else {
        v24 = (("foo bar" + ("hello" + v24)) + v24)
    }
    var v39 string = "foo bar"
    for fi40 := 0; fi40 < 5; fi40++ {
        _ = len(v3)
        for fi41 := 0; fi41 < 14; fi41++ {
            v34 -= v34
            _ = fi41
        }
        _ = fi40
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        v42 := GenS3{F0: 255, F1: true, F2: 65535}
        if true {
            for fi43 := 0; fi43 < 18; fi43++ {
                fmt.Println(v2, v34)
                v44 := v3[genFn1(-128, 1)]
                v45 := GenS4{F0: 42, F1: -32768}
                _ = fi43
                _ = v44
                _ = v45
            }
            for fi46 := 0; fi46 < 20; fi46++ {
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
                _ = fi46
            }
        } else {
            if false {
                v2 = v39
                fmt.Println(v39, v24)
            }
        }
        if false {
            for fi47 := 0; fi47 < 6; fi47++ {
                for fi48 := 0; fi48 < 17; fi48++ {
                    var v49 int8 = -128
                    v24 = (v39 + (v24 + "a"))
                    _ = fi48
                    _ = v49
                }
                _ = fi47
            }
            v50 := GenS5{F0: 853}
            _ = v50
        } else {
            if false {
                fmt.Println(v2, v39)
                v51 := GenS6{F0: v2, F1: 42, F2: 411}
                var v52 int16 = -486
                _ = v51
                _ = v52
            } else {
                var v53 int8 = 100
                if (585 != 4294967295) {
                    var v54 uint = 42
                    for fi55 := 0; fi55 < 8; fi55++ {
                        _ = len(v1)
                        v2 = v24
                        fmt.Println(v39)
                        _ = fi55
                    }
                    _ = v54
                } else {
                    v56 := GenS7{F0: 2147483647, F1: "a"}
                    _ = v56
                }
                var v57 uint32 = 4294967295
                _ = v53
                _ = v57
            }
        }
        _ = v42
    }
    for fi58 := 0; fi58 < 2; fi58++ {
        if (0 < 182) {
            var v59 int16 = int16(v35)
            switch v2 {
            case "":
                fmt.Println(v59, v24, fi58)
            }
            _ = v59
        }
        switch 2147483647 {
        case 32767:
            v35 -= (v35 << ((0) & 63))
        }
        _ = fi58
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v39
}
