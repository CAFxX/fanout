package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int8 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    v3 := []string{("hello" + ((("" + ("hello" + (("\n\t" + "foo bar") + "\n\t"))) + "\n\t") + "hello"))}
    v4 := GenS1{F0: 897, F1: "", F2: -122}
    switch -2147483648 {
    case 195:
        var v5 uint32 = 2147483647
        _ = v5
    case -1:
        var v6 float32 = float32(-2.25)
        _ = v6
    }
    v7 := GenS2{F0: 42}
    var v8 int64 = 0
    _ = len(v3)
    if (769 > 32768) {
        v8 += (v8 | 42)
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint64 = 935
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int8 = 57
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 string = "foo bar"
        v8 = v8
        v8 = 24
        _ = fi13
        _ = v14
    }
    switch 127 {
    case -9223372036854775808:
        v8 = (-(v8))
    case -871:
        var v15 int = 128
        _ = v15
    case -71:
        switch "a" {
        case "":
            var v16 bool = false
            _ = v16
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    if true {
        if (255 == 42) {
            v17 := GenS2{F0: -24}
            for fi18 := 0; fi18 < 10; fi18++ {
                for fi19 := 0; fi19 < 7; fi19++ {
                    var v20 int16 = 0
                    _ = fi19
                    _ = v20
                }
                if true {
                    v21 := GenS1{F0: 42, F1: ("a" + "foo bar"), F2: 100}
                    if (true || (!false)) {
                        var v22 uint64 = 514
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v22
                    } else {
                        if false {
                            v23 := GenS2{F0: 42}
                            v8 = v8
                            var v24 uint = 4294967295
                            _ = v23
                            _ = v24
                        }
                        var v25 string = "foo bar"
                        _ = v25
                    }
                    fmt.Println(fi18, v8)
                    _ = v21
                }
                _ = fi18
            }
            var v26 uint8 = 182
            _ = v17
            _ = v26
        }
        for fi27 := 0; fi27 < 3; fi27++ {
            v8 = v8
            var v28 uint = 759
            delete(v2, 893)
            _ = fi27
            _ = v28
        }
        if true {
            var v29 uint32 = 2147483648
            fmt.Println(v29, v29)
            var v30 int32 = 2147483647
            _ = v29
            _ = v30
        }
    } else {
        var v31 float32 = (0.0 - 0.0)
        for fi32 := 0; fi32 < 11; fi32++ {
            if true {
                fmt.Println(v8, v31)
                _ = len(v3)
                v8 = v8
            } else {
                v8 = -9223372036854775808
                v33 := v1[0]
                _ = v33
            }
            v34 := GenS2{F0: -8}
            _ = fi32
            _ = v34
        }
        var v35 int8 = -1
        _ = v31
        _ = v35
    }
    var v36 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v7
    _ = v8
}
