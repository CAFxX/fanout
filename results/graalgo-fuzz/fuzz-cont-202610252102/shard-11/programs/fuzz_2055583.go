package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 string; F2 int8 }

func main() {
    v1 := []uint32{2147483648, 65535}
    v2 := []int{128, 32768, -2147483648, -2147483648}
    var v3 uint = 18446744073709551615
    var v4 uint16 = 0
    v5 := make(map[string]int)
    var v6 float32 = 1e10
    v6 += float32(3.14159)
    v6 = (0.1 * v6)
    v7 := GenS1{F0: 520, F1: 0, F2: 4294967295}
    var v8 uint32 = 2
    v4 = 633
    switch "" {
    case "a":
        switch ((("" + "hello") + "") + "hello") {
        case "":
            v6 = ((v6 + float32(-2.25)) + 1e10)
        case "hello":
            delete(v5, "")
        case "a":
            v9 := v5[""]
            _ = v9
        default:
            if (!(v8 < v8)) {
                v10 := GenS1{F0: 127, F1: -2147483648, F2: 0}
                var v11 uint = v3
                _ = v10
                _ = v11
            } else {
                delete(v5, ("hello" + "\n\t"))
            }
        }
    }
    var v12 uint = (v3 % ((v3) | 1))
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS1{F0: fi13, F1: 253, F2: 0}
        _ = fi13
        _ = v14
    }
    if true {
        switch (("hello" + "foo bar") + "a") {
        case "":
            delete(v5, (("hello" + "") + (("a" + "\n\t") + (("" + "a") + ("\n\t" + (("a" + "foo bar") + "\n\t"))))))
        case "foo bar":
            fmt.Println(v6, v8)
        case "hello":
            v6 -= v6
        default:
            switch 1 {
            case -21:
                var v15 uint64 = 206
                _ = v15
            case 2147483647:
                if true {
                    v6 = v6
                }
            case -549:
                var v16 uint8 = 2
                _ = v16
            }
        }
        var v17 uint32 = uint32(v3)
        _ = v17
    }
    if ((true || false) || false) {
        if (!(1e10 >= 1e10)) {
            switch (("" + ("foo bar" + ("hello" + ("hello" + ("foo bar" + "foo bar"))))) + "") {
            case "a":
                v3 += v3
            case "":
                if (913 == 17) {
                    if false {
                        v18 := GenS2{F0: 42}
                        _ = v18
                    }
                    var v19 float32 = v6
                    _ = v19
                } else {
                    var v20 int32 = -1
                    _ = v20
                }
            default:
                v5[""] = 127
            }
            v8 = (v8 * v8)
        }
    }
    switch ("\n\t" + "a") {
    case "":
        v8 = (v8 * (v8 * (v8 << (((v3 << ((v12) & 63))) & 63))))
    }
    if false {
        var v21 uint32 = (v8 + v8)
        switch "\n\t" {
        case "a":
            for fi22 := 0; fi22 < 7; fi22++ {
                for fi23 := 0; fi23 < 9; fi23++ {
                    var v24 uint = v12
                    fmt.Println(v12, v24)
                    _ = fi23
                    _ = v24
                }
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 float64 = 1.5
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v8 += v8
                    _ = fi25
                    _ = v26
                }
                for fi27 := 0; fi27 < 15; fi27++ {
                    var v28 int32 = 522
                    _ = fi27
                    _ = v28
                }
                _ = fi22
            }
        case "hello":
            v6 = 0.1
        }
        _ = v21
    }
    v1 = append(v1, (v8 + (v8 % ((624) | 1))))
    switch "foo bar" {
    case "a":
        if len(v2) > 0 {
            v2[0] = 42
        }
    }
    var v29 string = "hello"
    if true {
        if false {
            v5[("a" + "\n\t")] = -468
            for fi30 := 0; fi30 < 15; fi30++ {
                v3 -= v3
                _ = fi30
            }
        }
    } else {
        v31 := GenS3{F0: (true || true), F1: ("foo bar" + v29), F2: 38}
        if false {
            v32 := GenS2{F0: 426}
            v29 = v29
            var v33 uint16 = v4
            _ = v32
            _ = v33
        }
        _ = v31
    }
    fmt.Println(v6, v6)
    switch "" {
    case "":
        fmt.Println(v4, v4, v8)
    case "a":
        var v34 float64 = (0.0 * float64(v6))
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
