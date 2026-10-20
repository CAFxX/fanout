package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }
func genFn0() uint64 {
    if (!true) {
        if false {
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 string = "\n\t"
                var v9 uint = 42
                v8 = ("\n\t" + "\n\t")
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        switch 295 {
        case 878:
            var v10 int = 888
            _ = v10
        }
    } else {
        var v11 uint64 = 210
        if true {
            var v12 int64 = -9223372036854775808
            fmt.Println(v11, v12)
            _ = v12
        } else {
            var v13 uint32 = 65535
            v14 := []int{42, 9223372036854775807, 4294967295, -9223372036854775808}
            v15 := GenS1{F0: 2147483647, F1: 771, F2: 133}
            _ = v13
            _ = v14
            _ = v15
        }
        fmt.Println(v11)
        _ = v11
    }
    if false {
        var v16 int64 = -302
        if (1.5 >= (1e10 * 3.14159)) {
            var v17 float32 = (3.14159 + 1.5)
            fmt.Println(v17, v17)
            _ = v17
        }
        v16 -= (-(v16))
        _ = v16
    }
    var v18 float64 = 3.14159
    if (!(0 >= 2147483647)) {
        v19 := GenS1{F0: 42, F1: 111, F2: 42}
        v18 = 1.5
        for fi20 := 0; fi20 < 13; fi20++ {
            fmt.Println(v18, fi20)
            var v21 uint = 65535
            v22 := []int{fi20, fi20, 0}
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = v19
    } else {
        fmt.Println(v18)
    }
    _ = v18
    return 2
}

func main() {
    var v1 uint64 = 2
    v2 := []string{""}
    v3 := []int{9223372036854775807, -2147483648, 128}
    v4 := []string{("hello" + "a"), ""}
    v5 := make(map[uint32]bool)
    var v6 int16 = 970
    var v23 bool = true
    v24 := GenS1{F0: 2, F1: 0, F2: 1}
    switch "\n\t" {
    case "\n\t":
        v6 -= v6
    case "hello":
        if len(v4) > 0 {
            v4[0] = ("" + ("hello" + "foo bar"))
        }
    default:
        if false {
            var v25 float64 = 3.14159
            v26 := GenS1{F0: 1, F1: 973, F2: 4294967295}
            _ = len(v5)
            _ = v25
            _ = v26
        } else {
            var v27 float32 = (0.0 + 0.0)
            v1 = v1
            var v28 uint16 = 975
            _ = v27
            _ = v28
        }
    }
    v29 := GenS1{F0: 728, F1: 2, F2: 4294967295}
    v30 := GenS1{F0: 83, F1: 1, F2: 632}
    var v31 uint = uint(v1)
    var v32 int8 = -1
    fmt.Println(v23, v32, v6)
    var v33 uint8 = 42
    var v34 uint64 = v1
    v35 := GenS1{F0: 72, F1: (v31 * uint(v33)), F2: 0}
    v34 += v1
    fmt.Println(v1, v6, v34)
    var v36 int64 = -82
    fmt.Println(v6)
    v33 += 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
