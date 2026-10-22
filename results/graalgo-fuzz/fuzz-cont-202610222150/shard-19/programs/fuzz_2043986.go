package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint64; F2 int }
type GenS4 struct { F0 int8; F1 uint16; F2 int8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint64, p1 int16, p2 uint8) int64 {
    var v6 int64 = -898
    var v7 int8 = 100
    switch ("a" + ("\n\t" + "\n\t")) {
    case "":
        if (639 > 775) {
            var v8 int32 = -2147483648
            var v9 uint32 = 626
            if (!false) {
                if false {
                    v10 := GenS1{F0: (^(p0))}
                    v11 := GenS1{F0: (p0 >> ((2) & 63))}
                    fmt.Println(v9)
                    _ = v10
                    _ = v11
                }
                if (true && (249 != 42)) {
                    v12 := GenS1{F0: 42}
                    fmt.Println(p0, v8, v6)
                    _ = v12
                }
                v13 := []int{-1, -976, 127, -676}
                _ = v13
            }
            _ = v8
            _ = v9
        }
    case "foo bar":
        fmt.Println(p2, v7)
    case "hello":
        for fi14 := 0; fi14 < 1; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 int16 = p1
                v16 = p1
                fmt.Println(fi14)
                p0 -= (p0 | (p0 & (^(p0))))
                p1 = p1
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    default:
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 float64 = (1.5 * 0.1)
            p1 = (p1 ^ -1)
            v19 := GenS1{F0: (^(p0))}
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    p1 = 0
    p0 += (p0 * 1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint64, p1 bool, p2 uint16) uint32 {
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int32 = 42
        fmt.Println(v21, fi20)
        for fi22 := 0; fi22 < 20; fi22++ {
            fmt.Println(p2)
            p1 = (p1 && p1)
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    switch -9223372036854775808 {
    case 675:
        var v23 int16 = -216
        _ = v23
    case -2147483648:
        var v24 uint16 = 146
        _ = v24
    default:
        v25 := []int{-1, -1, 42}
        _ = v25
    }
    var v26 int8 = -7
    for fi27 := 0; fi27 < 10; fi27++ {
        if (128 <= 128) {
            var v28 bool = ((-9223372036854775808 <= fi27) && (p1 || (v26 != (v26 / ((v26) | 1)))))
            v29 := []int{1, fi27, fi27}
            _ = v28
            _ = v29
        }
        p1 = true
        _ = fi27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return 362
}

func main() {
    var v1 float32 = 1e10
    var v2 string = "foo bar"
    v3 := make(map[string]string)
    var v4 uint8 = 127
    var v5 uint16 = 947
    var v30 int8 = 100
    switch "" {
    case "":
        fmt.Println(v4, v5, v30)
    case "hello":
        _ = len(v3)
    case "a":
        var v31 string = ("\n\t" + v2)
        _ = v31
    default:
        v32 := []int{42, 32767, 127, 1}
        _ = v32
    }
    switch -2147483648 {
    case 127:
        v33 := GenS2{F0: 465, F1: (v4 | (v4 ^ v4))}
        _ = v33
    case -9223372036854775808:
        fmt.Println(v4, v4, v4)
    }
    if (255 == (v5 & (^(v5)))) {
        v34 := []int{2147483647, 128, -807}
        switch ("hello" + ("hello" + v2)) {
        case "\n\t":
            v35 := GenS2{F0: 1, F1: v4}
            _ = v35
        default:
            v36 := GenS2{F0: genFn1(2, false, (^(v5))), F1: (v4 | (v4 | v4))}
            _ = v36
        }
        var v37 int32 = -564
        _ = v34
        _ = v37
    }
    delete(v3, v2)
    v38 := v3[((((("" + (("foo bar" + v2) + (((("a" + "a") + ((v2 + v2) + v2)) + (("foo bar" + v2) + "foo bar")) + (v2 + ((v2 + (("a" + (v2 + "\n\t")) + ("\n\t" + "\n\t"))) + v2))))) + ("hello" + v2)) + v2) + (v2 + (v2 + v2))) + v2)]
    var v39 float64 = ((1.5 / ((1.5 + 1.5) + 0.0)) + -2.25)
    var v40 int16 = 1000
    fmt.Println(v40, v30)
    if false {
        v30 -= v30
        fmt.Println(v30, v38)
    } else {
        v2 = (v38 + v38)
        for fi41 := 0; fi41 < 15; fi41++ {
            fmt.Println(v4, v2)
            v42 := GenS3{F0: 0, F1: 42, F2: fi41}
            _ = fi41
            _ = v42
        }
    }
    var v43 bool = (v30 >= 42)
    for fi44 := 0; fi44 < 20; fi44++ {
        if v43 {
            v45 := GenS4{F0: v30, F1: 485, F2: v30}
            v46 := GenS5{F0: v30}
            _ = v45
            _ = v46
        }
        var v47 int = fi44
        v48 := v3[v2]
        _ = fi44
        _ = v47
        _ = v48
    }
    switch v2 {
    case "hello":
        delete(v3, (v2 + ((v38 + ("a" + v38)) + v2)))
    }
    switch -2147483648 {
    case -961:
        var v49 int64 = 808
        _ = v49
    case 4294967295:
        var v50 string = ((v38 + "hello") + v38)
        _ = v50
    }
    var v51 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v5
    _ = v51
}
