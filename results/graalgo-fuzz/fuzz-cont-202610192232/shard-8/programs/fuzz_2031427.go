package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int16; F1 int32; F2 bool }

func main() {
    v1 := []uint32{4294967295, 2, 65535, 10}
    v2 := make(map[string]bool)
    v3 := []int{381, 32768, 32767}
    var v4 int16 = 32767
    var v5 int64 = 52
    v6 := []string{}
    var v7 float32 = 1.5
    var v8 uint64 = 2
    var v9 uint64 = v8
    v10 := GenS1{F0: 255, F1: 32767}
    for fi11 := 0; fi11 < 16; fi11++ {
        if true {
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 int32 = -2147483648
                switch 0 {
                case 127:
                    if (!false) {
                        fmt.Println(fi11, fi11, v13)
                        switch "\n\t" {
                        case "foo bar":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case "hello":
                            fmt.Println(v7)
                        }
                    }
                default:
                    var v14 uint8 = 233
                    _ = v14
                }
                _ = fi12
                _ = v13
            }
        } else {
            v9 += v8
        }
        v15 := GenS1{F0: 32767, F1: 0}
        var v16 uint32 = 734
        _ = fi11
        _ = v15
        _ = v16
    }
    switch 4294967295 {
    case 0:
        v7 = v7
    case 1:
        if len(v3) > 0 {
            v3[0] = -1
        }
    case -1:
        var v17 uint64 = v8
        _ = v17
    }
    var v18 int32 = int32(v5)
    var v19 int8 = 75
    switch 1 {
    case 42:
        v20 := GenS2{F0: 1}
        _ = v20
    case 128:
        if ((v9 >> ((1) & 63)) > (v8 ^ (v9 & v9))) {
            v21 := GenS2{F0: 4294967295}
            switch "a" {
            case "foo bar":
                for fi22 := 0; fi22 < 4; fi22++ {
                    if (v4 >= (^(v4))) {
                        var v23 uint32 = 1
                        var v24 int = -971
                        _ = v23
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 9; fi25++ {
                            v2["a"] = false
                            var v26 int16 = v4
                            _ = fi25
                            _ = v26
                        }
                        var v27 bool = (fi22 < fi22)
                        _ = v27
                    }
                    _ = fi22
                }
            default:
                fmt.Println(v4, v9, v7)
            }
            v8 -= (v9 << ((355) & 63))
            _ = v21
        }
    }
    switch (("a" + "hello") + "foo bar") {
    case "foo bar":
        switch 0 {
        case 32767:
            fmt.Println(v8, v9)
        default:
            v4 += int16(v18)
        }
    case "hello":
        v28 := GenS3{F0: v19, F1: "a"}
        _ = v28
    case "":
        v29 := GenS4{F0: 32767, F1: v18, F2: (true || true)}
        _ = v29
    default:
        switch (("a" + "hello") + ("a" + "foo bar")) {
        case "a":
            v30 := GenS4{F0: -32768, F1: (v18 >> ((4294967295) & 63)), F2: (-2.25 != 1.5)}
            _ = v30
        case "foo bar":
            var v31 uint8 = 42
            _ = v31
        case "":
            v32 := GenS1{F0: 255, F1: 32767}
            _ = v32
        default:
            v9 = (v8 & (^(v9)))
        }
    }
    if (18446744073709551615 == (v9 << ((0) & 63))) {
        switch "hello" {
        case "\n\t":
            v9 = (v9 * v9)
        }
        v4 = (-(v4))
    }
    _ = len(v2)
    if (true && (65535 > 255)) {
        var v33 uint16 = 0
        v2[("a" + ("a" + "\n\t"))] = (!false)
        _ = v33
    } else {
        var v34 int16 = v4
        var v35 string = ("foo bar" + (("a" + "foo bar") + "\n\t"))
        _ = v34
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
