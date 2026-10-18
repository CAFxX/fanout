package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint, p1 int64) uint8 {
    v4 := []int{9223372036854775807, 42, -619}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    p0 = p0
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    v5 := GenS1{F0: 761}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 128
}
func genFn1(p0 uint, p1 int64) uint8 {
    v6 := []int{-1, -1, 127, 42}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint64 = 0
        v9 := GenS1{F0: fi7}
        if false {
            switch ("a" + "foo bar") {
            case "":
                fmt.Println(p0, v8)
            case "a":
                for fi10 := 0; fi10 < 8; fi10++ {
                    _ = len(v6)
                    _ = fi10
                }
            case "\n\t":
                p1 += 982
            }
        } else {
            v6 = append(v6, (fi7 >> ((0) & 63)))
            var v11 int = (fi7 ^ fi7)
            if false {
                var v12 uint32 = 2147483648
                _ = v12
            }
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    p1 += p1
    var v13 float32 = 1e10
    var v14 bool = (914 == 2)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v6
    return 139
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]int)
    var v3 uint8 = 0
    var v15 uint = 219
    var v16 uint64 = 18446744073709551615
    if (v1 || (!v1)) {
        switch "hello" {
        case "hello":
            var v17 uint = v15
            _ = v17
        }
    } else {
        v16 = (v16 >> ((v15) & 63))
        v3 -= uint8(v16)
    }
    var v18 uint8 = v3
    for fi19 := 0; fi19 < 2; fi19++ {
        v15 = (v15 + v15)
        _ = fi19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 17; fi21++ {
            v22 := []int{-303, (^(fi20)), (fi20 & (fi21 / ((fi21) | 1))), (fi21 / (((fi20 >> (((v15 - (v15 | v15))) & 63))) | 1))}
            switch "" {
            case "foo bar":
                switch "\n\t" {
                case "foo bar":
                    v23 := GenS1{F0: fi20}
                    _ = v23
                case "\n\t":
                    if len(v22) > 0 {
                        v22[0] = fi20
                    }
                }
            case "hello":
                v24 := GenS1{F0: fi21}
                _ = v24
            default:
                delete(v2, 711)
            }
            v25 := GenS1{F0: fi21}
            _ = fi21
            _ = v22
            _ = v25
        }
        fmt.Println(fi20, v15)
        _ = fi20
    }
    var v26 float64 = 1.5
    for fi27 := 0; fi27 < 10; fi27++ {
        switch ("\n\t" + "") {
        case "a":
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 int = (fi27 | (fi28 + fi28))
                _ = fi28
                _ = v29
            }
        case "foo bar":
            var v30 float64 = (1.5 - 1e10)
            _ = v30
        case "\n\t":
            v31 := []int{fi27, fi27, fi27, 1}
            _ = v31
        }
        var v32 int64 = -9223372036854775808
        _ = fi27
        _ = v32
    }
    if (582 > 2147483648) {
        v26 -= v26
        v16 = v16
        v2[10] = int(v16)
    } else {
        v18 = genFn1(v15, 9223372036854775807)
        if ((1e10 * -2.25) != 0.0) {
            for fi33 := 0; fi33 < 14; fi33++ {
                var v34 float64 = 1.5
                fmt.Println(v1)
                var v35 int8 = 1
                _ = fi33
                _ = v34
                _ = v35
            }
        }
        var v36 uint8 = genFn0(369, 9223372036854775807)
        _ = v36
    }
    v3 -= uint8(v15)
    var v37 uint32 = 2
    fmt.Println(v26, v15)
    if ((int32(v16) < -1) && v1) {
        var v38 float64 = v26
        var v39 int32 = 1
        _ = v38
        _ = v39
    }
    if (!(42 == v37)) {
        v40 := []int{32768, -550, 128}
        fmt.Println(v18, v1)
        _ = v40
    } else {
        v3 = v3
        for fi41 := 0; fi41 < 8; fi41++ {
            v42 := []int{2147483647}
            fmt.Println(v15, v16, v16)
            if (2 > v15) {
                if len(v42) > 0 {
                    v42[0] = (fi41 - 327)
                }
                for fi43 := 0; fi43 < 15; fi43++ {
                    v15 += v15
                    _ = fi43
                }
            }
            _ = fi41
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v37
}
