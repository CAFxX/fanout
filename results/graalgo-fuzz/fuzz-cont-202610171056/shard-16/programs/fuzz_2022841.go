package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{32767}
    var v3 float64 = 1e10
    v4 := make(map[uint32]string)
    v5 := make(map[int]bool)
    var v6 uint = 4294967295
    if true {
        v3 -= (v3 / -2.25)
        var v7 string = "a"
        for fi8 := 0; fi8 < 18; fi8++ {
            if true {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 uint64 = 0
                    _ = len(v2)
                    _ = fi9
                    _ = v10
                }
            }
            var v11 uint16 = 42
            _ = fi8
            _ = v11
        }
        _ = v7
    }
    var v12 uint = 255
    fmt.Println(v3, v6, v12)
    switch ("a" + "a") {
    case "hello":
        switch -2147483648 {
        case 2147483647:
            var v13 uint16 = 255
            _ = v13
        }
    case "":
        if false {
            var v14 int8 = -52
            var v15 int16 = 32767
            _ = v14
            _ = v15
        }
    default:
        for fi16 := 0; fi16 < 1; fi16++ {
            _ = len(v2)
            _ = fi16
        }
    }
    var v17 int64 = 1
    var v18 bool = ("hello" >= (("" + "\n\t") + ("\n\t" + "hello")))
    v6 += (v12 & (v6 % (((v12 - v6)) | 1)))
    for fi19 := 0; fi19 < 14; fi19++ {
        if ("" < "a") {
            v20 := GenS1{F0: 604}
            v2 = append(v2, (^(fi19)))
            v21 := GenS1{F0: -32768}
            _ = v20
            _ = v21
        }
        var v22 uint64 = 2
        var v23 int32 = -2147483648
        _ = fi19
        _ = v22
        _ = v23
    }
    switch 9223372036854775807 {
    case 2147483647:
        switch "a" {
        case "foo bar":
            v3 = (v3 / float64(float32(v3)))
        case "\n\t":
            if false {
                for fi24 := 0; fi24 < 2; fi24++ {
                    switch int(v6) {
                    case 837:
                        v3 = v3
                    case 42:
                        var v25 uint64 = 4294967295
                        _ = v25
                    default:
                        v26 := v5[fi24]
                        _ = v26
                    }
                    var v27 int64 = v17
                    v28 := v4[990]
                    _ = fi24
                    _ = v27
                    _ = v28
                }
            } else {
                switch ("\n\t" + ("" + "foo bar")) {
                case "foo bar":
                    _ = len(v5)
                }
            }
        default:
            v29 := GenS1{F0: 42}
            _ = v29
        }
    default:
        v12 += v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
