package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int; F1 bool }
func genFn0() uint8 {
    v5 := []int{}
    if len(v5) > 0 {
        v5[0] = 1
    }
    _ = v5
    return 1
}

func main() {
    v1 := []uint32{2147483647, 65535, 65535, 4294967295}
    v2 := make(map[string]string)
    var v3 int64 = -9223372036854775808
    var v4 int8 = -1
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint32 = 751
        _ = fi6
        _ = v7
    }
    var v8 int32 = 1
    fmt.Println(v8, v3, v4)
    fmt.Println(v3, v3)
    if (false || true) {
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := GenS1{F0: v4, F1: 162, F2: 80}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 uint8 = 137
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 uint16 = uint16(v12)
                v2["a"] = "a"
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    }
    var v15 float32 = 0.0
    _ = len(v2)
    var v16 uint = 153
    if (!true) {
        v3 = v3
    }
    v17 := GenS2{F0: false, F1: v16, F2: 2}
    switch 793 {
    case -1:
        var v18 bool = (!false)
        _ = v18
    case 9223372036854775807:
        if false {
            v19 := GenS1{F0: v4, F1: 161, F2: 65535}
            if (-357 >= (v3 - 42)) {
                for fi20 := 0; fi20 < 11; fi20++ {
                    var v21 string = (("\n\t" + "foo bar") + "a")
                    var v22 int = (fi20 >> (((v16 * v16)) & 63))
                    switch ((("a" + v21) + v21) + (v21 + "a")) {
                    case "":
                        if (false || (!((3.14159 * 1.5) >= 0.1))) {
                            fmt.Println(fi20)
                        }
                    }
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                v23 := GenS3{F0: -32768, F1: 2, F2: 2}
                switch 78 {
                case -9223372036854775808:
                    if true {
                        v24 := GenS1{F0: v4, F1: 104, F2: 0}
                        for fi25 := 0; fi25 < 6; fi25++ {
                            var v26 int64 = 0
                            _ = fi25
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 11; fi27++ {
                            for fi28 := 0; fi28 < 8; fi28++ {
                                v29 := GenS2{F0: false, F1: 65535, F2: 2147483648}
                                _ = fi28
                                _ = v29
                            }
                            v30 := GenS4{F0: 1, F1: true}
                            _ = fi27
                            _ = v30
                        }
                        _ = v24
                    }
                }
                _ = v23
            }
            _ = v19
        } else {
            var v31 int16 = -844
            fmt.Println(v15, v3, v15)
            _ = v31
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
