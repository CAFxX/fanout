package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int; F1 uint; F2 int }
type GenS4 struct { F0 int; F1 bool; F2 uint32 }
type GenS5 struct { F0 string; F1 uint32; F2 int64 }
func genFn0(p0 int32, p1 int64) int32 {
    p1 += (p1 * (p1 & (-(p1))))
    p1 -= -1
    var v7 int = 840
    _ = p0
    _ = p1
    _ = v7
    return p0
}
func genFn1() int64 {
    var v8 uint16 = 65535
    var v9 float32 = 0.0
    v10 := GenS1{F0: 370, F1: -791}
    fmt.Println(v9)
    v9 = 0.0
    _ = v10
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint32 = 0
    var v2 int32 = 2147483647
    v3 := make(map[int]string)
    var v4 int = 799
    var v5 string = ((("a" + "\n\t") + "hello") + "a")
    var v6 uint8 = 42
    if true {
        for fi11 := 0; fi11 < 3; fi11++ {
            if true {
                fmt.Println(v6, fi11)
                v1 = v1
                v2 += genFn0(genFn0((v2 + v2), 0), 0)
            } else {
                v6 += (v6 + 2)
                v12 := GenS1{F0: genFn1(), F1: (v2 | (-(v2)))}
                _ = v12
            }
            fmt.Println(v4)
            var v13 uint64 = 1
            _ = fi11
            _ = v13
        }
        fmt.Println(v6, v4)
        v14 := []int{v4, v4, (v4 / (((v4 - v4)) | 1))}
        _ = v14
    }
    _ = len(v3)
    for fi15 := 0; fi15 < 13; fi15++ {
        fmt.Println(v2, v1)
        if true {
            fmt.Println(v2, v5, v1)
            fmt.Println(v4, v5, fi15)
            delete(v3, v4)
        } else {
            fmt.Println(v5, v2, v6)
        }
        var v16 uint32 = v1
        _ = fi15
        _ = v16
    }
    v2 = 42
    if (4294967295 != v4) {
        v17 := GenS1{F0: -9223372036854775808, F1: 43}
        var v18 uint8 = v6
        switch v4 {
        case 127:
            v19 := GenS1{F0: genFn1(), F1: v2}
            _ = v19
        case -1:
            var v20 int = (v4 & v4)
            _ = v20
        }
        _ = v17
        _ = v18
    } else {
        switch ("a" + (v5 + ("foo bar" + v5))) {
        case "foo bar":
            v21 := GenS1{F0: -314, F1: v2}
            _ = v21
        default:
            v2 = (v2 * 934)
        }
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        for fi23 := 0; fi23 < 16; fi23++ {
            if false {
                v2 = -1
                var v24 int32 = (v2 + (v2 & v2))
                fmt.Println(v24, v24, fi23)
                _ = v24
            }
            v5 = "\n\t"
            switch "\n\t" {
            case "":
                fmt.Println(fi23, v1, v1)
            case "a":
                switch ((v5 + (v5 + "")) + "") {
                case "hello":
                    v1 = (v1 * (^(v1)))
                case "a":
                    var v25 int64 = -1
                    v26 := GenS1{F0: v25, F1: genFn0((v2 - 1), v25)}
                    _ = v25
                    _ = v26
                case "":
                    fmt.Println(v6)
                }
                if (int64(v6) >= 0) {
                    v27 := GenS1{F0: 9223372036854775807, F1: int32(v1)}
                    v28 := GenS2{F0: 4294967295, F1: -1}
                    var v29 uint8 = (v6 | (v6 ^ 127))
                    _ = v27
                    _ = v28
                    _ = v29
                }
                fmt.Println(v1)
            }
            _ = fi23
        }
        if (((fi22 / ((32768) | 1)) == (fi22 >> ((123) & 63))) || (true || true)) {
            var v30 int8 = 0
            _ = v30
        }
        v31 := GenS3{F0: (v4 << ((459) & 63)), F1: 1, F2: v4}
        _ = fi22
        _ = v31
    }
    v32 := GenS2{F0: 18446744073709551615, F1: -128}
    v33 := []int{}
    v34 := GenS2{F0: 42, F1: -124}
    var v35 uint8 = v6
    if len(v33) >= 2 {
        v33 = v33[1:len(v33)]
    }
    v35 = 187
    if (4294967295 <= 95) {
        v35 -= v35
    } else {
        v36 := GenS4{F0: v4, F1: false, F2: 0}
        var v37 float64 = 0.0
        var v38 int64 = -9223372036854775808
        _ = v36
        _ = v37
        _ = v38
    }
    switch (v4 - -649) {
    case 2147483647:
        fmt.Println(v5)
    case 9223372036854775807:
        var v39 uint8 = 2
        _ = v39
    case 127:
        fmt.Println(v6, v35, v4)
    default:
        v40 := GenS5{F0: "\n\t", F1: (v1 << ((uint(v6)) & 63)), F2: -9223372036854775808}
        _ = v40
    }
    fmt.Println(v35, v1)
    for fi41 := 0; fi41 < 2; fi41++ {
        if false {
            var v42 bool = false
            _ = v42
        } else {
            if false {
                var v43 uint = 645
                _ = v43
            }
            fmt.Println(v2)
            fmt.Println(fi41, v1)
        }
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
