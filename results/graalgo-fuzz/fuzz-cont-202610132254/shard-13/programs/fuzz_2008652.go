package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint }
func genFn0(p0 int) int32 {
    p0 += (p0 & 4294967295)
    switch (p0 >> ((640) & 63)) {
    case 631:
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 int8 = 35
            _ = fi7
            _ = v8
        }
    default:
        v9 := GenS1{F0: true, F1: 130}
        _ = v9
    }
    p0 += p0
    _ = p0
    return -1
}
func genFn1(p0 uint, p1 string) uint64 {
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := []int{(^(fi10))}
        v12 := GenS2{F0: 255, F1: 65535}
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 42, F1: (p0 * 790)}
    v14 := []int{}
    p1 = ""
    p0 = p0
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 uint8 = 255
    var v5 int32 = 2147483647
    v6 := make(map[string]bool)
    var v15 int32 = (v5 | v5)
    var v16 uint = 959
    v1["a"] = 0
    switch "" {
    case "":
        if true {
            switch 42 {
            case -200:
                v5 -= genFn0(42)
            default:
                v15 = genFn0(128)
            }
            var v17 uint64 = 2
            var v18 uint64 = (v17 | 0)
            _ = v17
            _ = v18
        }
    case "\n\t":
        v19 := GenS2{F0: 2, F1: 65535}
        _ = v19
    case "hello":
        switch "hello" {
        case "":
            var v20 uint32 = 929
            _ = v20
        case "foo bar":
            v21 := []int{}
            _ = v21
        }
    default:
        var v22 uint8 = 59
        _ = v22
    }
    var v23 uint = v16
    for fi24 := 0; fi24 < 3; fi24++ {
        if false {
            v23 += 2
        }
        v16 = (v23 | v16)
        _ = fi24
    }
    v5 = (-(v15))
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 uint8 = uint8(v16)
        v15 = v15
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 bool = ((false || (genFn0(fi27) < v5)) || ((-1 > 42) && (42 < uint(v4))))
        var v29 string = ("a" + "a")
        if (v16 >= v16) {
            v5 -= v15
            v30 := []int{}
            _ = v30
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    v5 = v15
    var v31 float32 = (float32(0.0) + 3.14159)
    for fi32 := 0; fi32 < 15; fi32++ {
        switch (fi32 ^ fi32) {
        case 32767:
            var v33 string = "hello"
            _ = v33
        case 4294967295:
            fmt.Println(v15, v5, fi32)
        case -281:
            var v34 int = -2147483648
            _ = v34
        default:
            for fi35 := 0; fi35 < 13; fi35++ {
                fmt.Println(v31, fi35)
                _ = fi35
            }
        }
        switch "" {
        case "a":
            v36 := v2[("foo bar" + "a")]
            _ = v36
        case "\n\t":
            for fi37 := 0; fi37 < 7; fi37++ {
                var v38 float32 = 3.14159
                switch ((("" + ("\n\t" + "foo bar")) + "a") + "a") {
                case "a":
                    switch (fi32 ^ 1) {
                    case 2147483647:
                        v39 := GenS2{F0: (v4 | (v4 & v4)), F1: 2}
                        _ = v39
                    case 0:
                        v40 := []int{}
                        _ = v40
                    }
                case "hello":
                    var v41 bool = false
                    _ = v41
                case "foo bar":
                    var v42 bool = ((^(v23)) > 0)
                    _ = v42
                }
                _ = fi37
                _ = v38
            }
        default:
            v43 := []int{4294967295, fi32, fi32}
            _ = v43
        }
        _ = fi32
    }
    var v44 uint = v16
    if false {
        v31 += v31
    } else {
        var v45 uint64 = 4294967295
        v46 := GenS2{F0: v4, F1: 4294967295}
        var v47 uint16 = 0
        _ = v45
        _ = v46
        _ = v47
    }
    var v48 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v4
    _ = v44
    _ = v48
    _ = v5
    _ = v6
}
