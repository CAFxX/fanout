package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint; F2 string }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint64) int64 {
    v6 := GenS1{F0: (p0 >> ((523) & 63)), F1: p0}
    switch 4294967295 {
    case 509:
        v7 := []int{0, 4294967295}
        _ = v7
    case -403:
        switch (("" + "") + "hello") {
        case "foo bar":
            v8 := GenS2{F0: 428, F1: 0, F2: "hello"}
            _ = v8
        case "hello":
            v9 := GenS2{F0: 65535, F1: 482, F2: ("\n\t" + ("" + ""))}
            _ = v9
        }
    }
    p0 = 18446744073709551615
    var v10 string = ("\n\t" + "foo bar")
    var v11 uint64 = 227
    _ = p0
    _ = v10
    _ = v11
    _ = v6
    return 0
}

func main() {
    var v1 int16 = 1
    var v2 uint64 = 0
    var v3 string = ""
    var v4 int = -2147483648
    v5 := []int{(v4 << ((478) & 63)), (v4 | v4), (v4 ^ v4)}
    if ((false && ((-3 == 127) || (!(false || false)))) && true) {
        if (2 < v2) {
            fmt.Println(v1)
        }
        var v12 uint64 = v2
        _ = v12
    }
    v3 = v3
    v1 = v1
    if (v1 >= v1) {
        v1 += v1
        if true {
            switch ((v3 + ((("a" + v3) + "hello") + "hello")) + "foo bar") {
            case "hello":
                for fi13 := 0; fi13 < 17; fi13++ {
                    var v14 float32 = ((3.14159 / 1.0) / -2.25)
                    v2 -= 727
                    _ = fi13
                    _ = v14
                }
            case "\n\t":
                if (genFn0((v2 ^ v2)) > 1) {
                    var v15 int = -561
                    _ = v15
                }
            case "foo bar":
                v4 += 1
            default:
                v4 = v4
            }
        } else {
            v2 += 1
            var v16 uint = 227
            _ = v16
        }
    } else {
        v17 := GenS2{F0: 781, F1: 18446744073709551615, F2: v3}
        v1 -= 42
        _ = v17
    }
    v18 := GenS2{F0: 65535, F1: 2, F2: "\n\t"}
    v19 := GenS2{F0: 1, F1: 65535, F2: "hello"}
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS1{F0: (^(v2)), F1: v2}
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 float64 = 1e10
            v3 = "\n\t"
            v23 = 0.1
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    if len(v5) > 0 {
        v5[0] = (v4 & v4)
    }
    var v24 float32 = float32((3.14159 + 0.0))
    _ = len(v5)
    for fi25 := 0; fi25 < 14; fi25++ {
        var v26 uint64 = v2
        switch (v4 - (v4 * v4)) {
        case 127:
            v27 := GenS2{F0: 65535, F1: 1, F2: (((v3 + v3) + v3) + v3)}
            _ = v27
        case 128:
            var v28 uint64 = 42
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    v5 = append(v5, v4)
    if (354 > 1) {
        var v29 uint8 = 127
        if (!true) {
            v30 := GenS2{F0: 42, F1: 1, F2: v3}
            _ = v30
        }
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 int64 = 0
            _ = fi31
            _ = v32
        }
        _ = v29
    }
    var v33 int16 = v1
    switch v3 {
    case "hello":
        fmt.Println(v3, v4, v2)
    case "":
        for fi34 := 0; fi34 < 14; fi34++ {
            for fi35 := 0; fi35 < 14; fi35++ {
                v2 = (v2 % (((v2 * v2)) | 1))
                v36 := GenS3{F0: 42}
                v37 := GenS2{F0: 473, F1: 220, F2: ""}
                _ = fi35
                _ = v36
                _ = v37
            }
            if false {
                v2 -= v2
                var v38 int = v4
                _ = v38
            } else {
                v4 -= (v4 & fi34)
            }
            v39 := GenS2{F0: 573, F1: 0, F2: (v3 + v3)}
            _ = fi34
            _ = v39
        }
    }
    var v40 float64 = ((3.14159 - 0.0) + 0.1)
    fmt.Println(v40, v1)
    v5 = append(v5, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v40
    _ = v5
}
