package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint32; F2 int32 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 int16 = -970
    v2 := make(map[uint32]bool)
    v3 := []int{170, 127, 32767, 128}
    v4 := make(map[int]string)
    var v5 string = "a"
    v6 := v2[120]
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v10 uint32 = 0
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v11 int32 = -277
    var v12 int = -320
    v13 := GenS1{F0: 4294967295}
    switch v5 {
    case "a":
        switch v5 {
        case "hello":
            fmt.Println(v5, v1)
        case "foo bar":
            v12 = (v12 | (v12 / (((v12 | v12)) | 1)))
        default:
            fmt.Println(v1)
        }
    case "":
        var v14 int32 = v11
        _ = v14
    case "foo bar":
        var v15 int16 = 32767
        _ = v15
    }
    if v6 {
        if false {
            v1 = (v1 % ((v1) | 1))
            for fi16 := 0; fi16 < 2; fi16++ {
                if v6 {
                    v17 := GenS1{F0: 4294967295}
                    v18 := GenS1{F0: 42}
                    _ = v17
                    _ = v18
                }
                _ = fi16
            }
        }
    }
    switch (^(v12)) {
    case 42:
        for fi19 := 0; fi19 < 19; fi19++ {
            v20 := GenS2{F0: -2147483648}
            v5 = (((("" + (v5 + v5)) + "") + (v5 + v5)) + v5)
            _ = fi19
            _ = v20
        }
    default:
        if len(v3) > 0 {
            v3[0] = (v12 << ((0) & 63))
        }
    }
    switch ("hello" + (v5 + "hello")) {
    case "\n\t":
        switch v12 {
        case 910:
            v1 = 830
        }
    case "hello":
        for fi21 := 0; fi21 < 12; fi21++ {
            v22 := GenS3{F0: -13, F1: 590, F2: v11}
            _ = fi21
            _ = v22
        }
    case "foo bar":
        for fi23 := 0; fi23 < 3; fi23++ {
            switch ((v5 + v5) + "\n\t") {
            case "hello":
                if true {
                    if (127 == 42) {
                        switch v5 {
                        case "hello":
                            var v24 uint64 = 18446744073709551615
                            _ = v24
                        }
                        var v25 int32 = -1
                        _ = v25
                    }
                    fmt.Println(fi23, v6, v11)
                    if (!(v6 && (42 >= 146))) {
                        v26 := GenS3{F0: -5, F1: 1, F2: 1}
                        for fi27 := 0; fi27 < 2; fi27++ {
                            v6 = v6
                            _ = fi27
                        }
                        v1 = 32767
                        _ = v26
                    }
                }
            case "foo bar":
                v28 := GenS4{F0: v1}
                _ = v28
            }
            v6 = v6
            _ = fi23
        }
    default:
        v5 = v5
    }
    v4[(-(v12))] = ""
    v5 = v5
    v3 = append(v3, 128)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
