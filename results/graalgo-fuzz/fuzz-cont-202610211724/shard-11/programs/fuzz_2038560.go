package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int64 }
type GenS2 struct { F0 int; F1 uint32; F2 int16 }
func genFn0(p0 int64, p1 uint, p2 bool) uint64 {
    v7 := GenS1{F0: "hello", F1: 255, F2: p0}
    var v8 int16 = 882
    var v9 uint16 = 924
    v8 += (v8 << ((2) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := []int{-731}
    v2 := make(map[int]string)
    v3 := []uint32{}
    v4 := make(map[uint32]string)
    var v5 bool = false
    var v6 int32 = -1
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := v4[65535]
            v12 = v12
            var v13 int8 = -33
            _ = fi11
            _ = v12
            _ = v13
        }
        if (!(!v5)) {
            if (42 < 65535) {
                v6 += v6
                switch "foo bar" {
                case "hello":
                    var v14 uint64 = 1
                    _ = v14
                case "a":
                    var v15 int16 = -152
                    _ = v15
                }
                var v16 float64 = 0.1
                _ = v16
            } else {
                v6 -= -2147483648
            }
            v5 = (255 < 758)
            var v17 int = (fi10 | (fi10 / ((fi10) | 1)))
            _ = v17
        } else {
            var v18 int32 = (v6 % ((0) | 1))
            switch fi10 {
            case -532:
                var v19 uint64 = 18446744073709551615
                _ = v19
            case 962:
                switch (("\n\t" + "\n\t") + "foo bar") {
                case "\n\t":
                    v20 := v2[-2147483648]
                    _ = v20
                }
            default:
                v18 = v6
            }
            v5 = (1.5 < 1e10)
            _ = v18
        }
        _ = fi10
    }
    if (42 < 1000) {
        var v21 uint16 = 42
        _ = v21
    }
    switch "a" {
    case "hello":
        v22 := v2[32768]
        _ = v22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        switch (("a" + "") + "\n\t") {
        case "foo bar":
            if v5 {
                var v24 uint = 4294967295
                _ = v24
            } else {
                fmt.Println(v5, v6)
            }
        case "":
            v25 := v2[(fi23 >> ((255) & 63))]
            _ = v25
        case "a":
            var v26 int64 = 42
            _ = v26
        default:
            v5 = v5
        }
        _ = fi23
    }
    if (v5 || false) {
        var v27 uint64 = genFn0(0, 65535, v5)
        v6 = (v6 % ((v6) | 1))
        v6 = (^(v6))
        _ = v27
    } else {
        switch "" {
        case "\n\t":
            var v28 uint = 255
            _ = v28
        default:
            var v29 int = -1
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        if (863 >= 724) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v6, fi30, v6)
        } else {
            if true {
                v6 = 1
            }
            v31 := GenS1{F0: "a", F1: 0, F2: 42}
            v32 := GenS1{F0: ("foo bar" + (("" + "foo bar") + (("\n\t" + ("hello" + "a")) + ""))), F1: 4294967295, F2: -28}
            _ = v31
            _ = v32
        }
        var v33 uint32 = 42
        _ = fi30
        _ = v33
    }
    var v34 int64 = -856
    for fi35 := 0; fi35 < 2; fi35++ {
        var v36 int64 = v34
        _ = fi35
        _ = v36
    }
    v4[65535] = "hello"
    if (v34 == v34) {
        var v37 int = -28
        _ = v37
    }
    v38 := GenS1{F0: ((("foo bar" + "\n\t") + "a") + ((("hello" + (("hello" + "hello") + "\n\t")) + ("\n\t" + ("hello" + "\n\t"))) + "")), F1: 2, F2: v34}
    for fi39 := 0; fi39 < 2; fi39++ {
        v40 := GenS1{F0: "hello", F1: 0, F2: 42}
        v34 -= v34
        _ = fi39
        _ = v40
    }
    v6 += 486
    v41 := GenS2{F0: 1, F1: 2147483647, F2: 32767}
    v5 = (v5 && v5)
    var v42 uint32 = 530
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
