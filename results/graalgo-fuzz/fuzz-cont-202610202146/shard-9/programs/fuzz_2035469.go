package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int32; F1 string; F2 int16 }
func genFn0() uint64 {
    if (355 != 1) {
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := GenS1{F0: 0}
            v9 := GenS2{F0: 0, F1: 936}
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        v10 := []int{-1}
        _ = v10
    }
    v11 := []int{}
    var v12 int8 = 127
    v12 = v12
    _ = v11
    _ = v12
    return 18446744073709551615
}
func genFn1(p0 int16, p1 bool, p2 uint8) uint8 {
    if false {
        if p1 {
            if p1 {
                v13 := GenS1{F0: 518}
                v14 := GenS3{F0: 4294967295, F1: 0}
                _ = v13
                _ = v14
            }
            fmt.Println(p0)
        }
    } else {
        switch ("" + "hello") {
        case "":
            var v15 int16 = 1000
            _ = v15
        case "\n\t":
            var v16 int8 = 100
            _ = v16
        case "hello":
            fmt.Println(p0, p0, p2)
        }
    }
    var v17 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 255
    v3 := make(map[uint32]int)
    var v4 int64 = 42
    var v5 uint = 0
    v6 := make(map[uint32]bool)
    fmt.Println(v2, v4)
    v5 = (v5 * (v2 << ((v5) & 63)))
    var v18 uint = (v5 / ((586) | 1))
    for fi19 := 0; fi19 < 15; fi19++ {
        fmt.Println(fi19, v18, v5)
        for fi20 := 0; fi20 < 3; fi20++ {
            v21 := GenS2{F0: v2, F1: -2147483648}
            var v22 float32 = -2.25
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    v5 = (^(v18))
    for fi23 := 0; fi23 < 14; fi23++ {
        v4 = -543
        for fi24 := 0; fi24 < 19; fi24++ {
            v4 = 238
            switch ("foo bar" + ("\n\t" + "foo bar")) {
            case "\n\t":
                var v25 int8 = 52
                v26 := GenS3{F0: 2147483647, F1: 0}
                v27 := v6[1]
                v28 := GenS3{F0: 2, F1: (^(v25))}
                v27 = v27
                _ = v25
                _ = v26
                _ = v27
                _ = v28
            }
            v29 := GenS4{F0: -920, F1: "foo bar", F2: 42}
            _ = fi24
            _ = v29
        }
        var v30 uint32 = 42
        _ = fi23
        _ = v30
    }
    switch ("a" + ("" + "\n\t")) {
    case "":
        fmt.Println(v4)
    case "foo bar":
        v4 = -905
    default:
        if (false || false) {
            v31 := GenS4{F0: -1, F1: ("hello" + ("foo bar" + "foo bar")), F2: 42}
            var v32 int16 = -589
            _ = v31
            _ = v32
        }
    }
    var v33 int64 = v4
    for fi34 := 0; fi34 < 14; fi34++ {
        fmt.Println(fi34, v4, v5)
        var v35 uint64 = 371
        v36 := GenS4{F0: 0, F1: "a", F2: -1}
        _ = fi34
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 20; fi37++ {
        for fi38 := 0; fi38 < 10; fi38++ {
            switch "foo bar" {
            case "hello":
                var v39 int = 139
                _ = v39
            case "foo bar":
                v40 := GenS2{F0: 577, F1: 42}
                _ = v40
            }
            _ = fi38
        }
        var v41 string = "foo bar"
        var v42 int = fi37
        _ = fi37
        _ = v41
        _ = v42
    }
    var v43 int32 = 1
    v44 := []int{}
    for fi45 := 0; fi45 < 14; fi45++ {
        switch fi45 {
        case -2147483648:
            if (!(float64((0.1 / 1e10)) > 0.1)) {
                v46 := GenS1{F0: 4294967295}
                _ = v46
            }
        }
        var v47 bool = false
        v33 -= v4
        _ = fi45
        _ = v47
    }
    fmt.Println(v18, v33, v2)
    var v48 float64 = (-2.25 / 1.5)
    fmt.Println(v18)
    var v49 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v43
    _ = v44
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
