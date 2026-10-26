package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int32; F2 int16 }
type GenS3 struct { F0 string }
func genFn0() int32 {
    var v6 uint16 = 32767
    fmt.Println(v6)
    switch 384 {
    case -516:
        switch -143 {
        case -747:
            var v7 uint = 0
            _ = v7
        default:
            v8 := GenS1{F0: 32767, F1: 4294967295}
            _ = v8
        }
    case -530:
        var v9 int = 128
        _ = v9
    default:
        var v10 int8 = -34
        _ = v10
    }
    v6 = 271
    _ = v6
    return 0
}

func main() {
    v1 := []string{("" + "")}
    v2 := []uint32{599, 1, 1, 2}
    var v3 int = -9223372036854775808
    var v4 uint16 = 42
    v5 := []int{v3}
    v3 += v3
    v11 := GenS1{F0: 42, F1: 0}
    if true {
        fmt.Println(v3, v3)
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint16 = (v4 >> ((658) & 63))
        if true {
            v13 = (v13 | v4)
            switch "" {
            case "foo bar":
                v14 := GenS2{F0: 149, F1: 0, F2: 112}
                _ = v14
            }
            v2 = append(v2, 4)
        }
        var v15 float32 = float32(1e10)
        _ = fi12
        _ = v13
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        fmt.Println(fi16, v4, fi16)
        switch "foo bar" {
        case "foo bar":
            v17 := GenS1{F0: (v3 & v3), F1: 813}
            _ = v17
        default:
            v18 := GenS2{F0: 42, F1: 2147483647, F2: 564}
            _ = v18
        }
        v5 = append(v5, 1)
        _ = fi16
    }
    v3 = v3
    var v19 int32 = 358
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = (^(v3))
    var v20 bool = false
    var v21 bool = v20
    if true {
        for fi22 := 0; fi22 < 17; fi22++ {
            var v23 int8 = 42
            if v21 {
                switch (-(v3)) {
                case 42:
                    switch ("hello" + "hello") {
                    case "\n\t":
                        var v24 uint16 = (v4 | 65535)
                        _ = v24
                    }
                case 665:
                    for fi25 := 0; fi25 < 11; fi25++ {
                        switch v3 {
                        case -9223372036854775808:
                            v19 -= genFn0()
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(v21, fi25, v3)
                        v26 := GenS3{F0: (("" + ("foo bar" + "hello")) + (("foo bar" + "hello") + ((("" + ("a" + ("\n\t" + ("hello" + ("foo bar" + "foo bar"))))) + (("foo bar" + (("foo bar" + "\n\t") + ("hello" + ("hello" + "")))) + "foo bar")) + ((("a" + "foo bar") + "foo bar") + "hello"))))}
                        _ = fi25
                        _ = v26
                    }
                case 32767:
                    v27 := GenS3{F0: "a"}
                    _ = v27
                }
                v28 := GenS3{F0: ("\n\t" + "a")}
                _ = v28
            } else {
                var v29 int16 = -439
                v21 = (2 > 0)
                _ = v29
            }
            fmt.Println(v19, v23, v21)
            _ = fi22
            _ = v23
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
