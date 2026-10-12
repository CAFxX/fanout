package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint32 }
func genFn0() uint {
    switch "foo bar" {
    case "foo bar":
        switch 128 {
        case 328:
            var v5 int64 = 0
            _ = v5
        }
    case "hello":
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := GenS1{F0: -1}
            v8 := GenS2{F0: 573}
            _ = fi6
            _ = v7
            _ = v8
        }
    case "":
        if ("\n\t" >= ("foo bar" + "")) {
            var v9 int8 = -57
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 uint16 = 0
                var v12 uint32 = 652
                var v13 int32 = 42
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            _ = v9
        }
    }
    var v14 bool = false
    var v15 uint16 = 876
    _ = v14
    _ = v15
    return 0
}

func main() {
    var v1 int8 = 127
    var v2 uint8 = 165
    v3 := []uint32{340, 2147483647, 2147483648, 0}
    v4 := make(map[string]bool)
    if true {
        var v16 uint32 = uint32(v2)
        switch ((("\n\t" + ("\n\t" + "hello")) + ((("foo bar" + ("foo bar" + "\n\t")) + ("foo bar" + "\n\t")) + "hello")) + "a") {
        case "hello":
            var v17 int64 = 1
            _ = v17
        default:
            fmt.Println(v16, v1)
        }
        if false {
            v18 := GenS1{F0: 2147483647}
            var v19 uint = 185
            v19 = uint(v2)
            _ = v18
            _ = v19
        } else {
            _ = len(v4)
        }
        _ = v16
    } else {
        v20 := GenS3{F0: 2147483647}
        _ = v20
    }
    switch -872 {
    case 4294967295:
        v2 -= (v2 | v2)
    case 32768:
        for fi21 := 0; fi21 < 7; fi21++ {
            v1 -= v1
            v22 := GenS4{F0: (v2 % (((v2 >> ((42) & 63))) | 1))}
            for fi23 := 0; fi23 < 4; fi23++ {
                fmt.Println(fi21, v1, fi21)
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
    }
    fmt.Println(v2, v2)
    v2 = (v2 / ((42) | 1))
    for fi24 := 0; fi24 < 12; fi24++ {
        v4[("foo bar" + (("foo bar" + ((("\n\t" + "\n\t") + "hello") + "foo bar")) + "a"))] = false
        _ = fi24
    }
    var v25 bool = false
    if len(v3) > 0 {
        v3[0] = 872
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        if (141 != 255) {
            var v27 string = ("hello" + "")
            _ = v27
        }
        for fi28 := 0; fi28 < 18; fi28++ {
            _ = len(v4)
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
            _ = fi28
        }
        switch ("\n\t" + "a") {
        case "hello":
            v1 += (v1 >> ((188) & 63))
        default:
            for fi29 := 0; fi29 < 19; fi29++ {
                var v30 uint32 = 2147483648
                switch (fi26 + fi26) {
                case 127:
                    if ((fi26 + -1) < 2147483647) {
                        switch "\n\t" {
                        case "a":
                            var v31 uint32 = v30
                            _ = v31
                        case "foo bar":
                            var v32 int = fi29
                            _ = v32
                        default:
                            var v33 uint64 = 242
                            _ = v33
                        }
                    } else {
                        v34 := GenS5{F0: ("" + "hello"), F1: 0}
                        var v35 int64 = 42
                        var v36 string = "a"
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                default:
                    v30 = (v30 / ((v30) | 1))
                }
                _ = fi29
                _ = v30
            }
        }
        _ = fi26
    }
    fmt.Println(v25, v25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
