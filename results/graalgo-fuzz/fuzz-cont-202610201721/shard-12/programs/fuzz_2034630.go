package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
func genFn0() uint {
    v6 := GenS1{F0: 691, F1: 4294967295, F2: 9223372036854775807}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint64 = 2
        v9 := GenS2{F0: -1}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v6
    return 0
}

func main() {
    var v1 uint = 2
    var v2 int64 = 896
    v3 := []string{"foo bar", (("" + ("hello" + ("a" + ((((((("\n\t" + "") + "foo bar") + "a") + "foo bar") + "\n\t") + ("foo bar" + "")) + "foo bar")))) + "")}
    v4 := make(map[int]uint32)
    var v5 int64 = v2
    var v10 uint16 = uint16(v1)
    var v11 int = 32767
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 string = ((("a" + "\n\t") + "hello") + "foo bar")
        switch ("a" + v13) {
        case "foo bar":
            v4[(v11 << (((v1 + v1)) & 63))] = 0
        default:
            fmt.Println(fi12, v5)
        }
        _ = fi12
        _ = v13
    }
    var v14 int16 = 18
    for fi15 := 0; fi15 < 7; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            switch (v11 - 127) {
            case -9223372036854775808:
                _ = len(v3)
            }
            v17 := GenS1{F0: v14, F1: 1, F2: fi15}
            _ = fi16
            _ = v17
        }
        var v18 bool = false
        v5 = (^(v5))
        _ = fi15
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        switch 763 {
        case 9223372036854775807:
            v20 := GenS3{F0: -2147483648}
            _ = v20
        case 128:
            var v21 int16 = (v14 & 87)
            _ = v21
        }
        for fi22 := 0; fi22 < 3; fi22++ {
            if true {
                switch ("foo bar" + "a") {
                case "\n\t":
                    v1 -= (v1 ^ (v1 % ((42) | 1)))
                case "":
                    v3 = append(v3, "foo bar")
                case "hello":
                    var v23 bool = false
                    _ = v23
                }
            }
            switch ("foo bar" + "\n\t") {
            case "hello":
                v2 = v2
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    if (!(!false)) {
                        var v24 float32 = (1e10 - -2.25)
                        _ = v24
                    } else {
                        v10 = uint16(v1)
                        v14 -= 1
                        fmt.Println(fi19)
                    }
                case "hello":
                    var v25 uint = genFn0()
                    _ = v25
                default:
                    v2 = (^(v5))
                }
            default:
                v3 = append(v3, ("foo bar" + ""))
                delete(v4, (v11 & fi19))
                v14 -= 0
            }
            _ = fi22
        }
        var v26 uint16 = v10
        _ = fi19
        _ = v26
    }
    v2 += v2
    v14 = 1000
    v27 := GenS2{F0: 1000}
    for fi28 := 0; fi28 < 15; fi28++ {
        _ = len(v3)
        _ = fi28
    }
    if true {
        _ = len(v4)
        switch v11 {
        case -382:
            var v29 uint16 = v10
            _ = v29
        case 9223372036854775807:
            v11 += v11
        }
        var v30 int32 = 0
        _ = v30
    } else {
        v31 := GenS3{F0: 662}
        _ = v31
    }
    switch (v11 * v11) {
    case 4294967295:
        var v32 uint64 = uint64(v1)
        _ = v32
    default:
        if ((v11 | (v11 | v11)) == v11) {
            fmt.Println(v1, v5)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v10 = v10
        } else {
            var v33 uint32 = 4294967295
            var v34 int64 = (v5 ^ (v2 >> ((genFn0()) & 63)))
            v1 += genFn0()
            _ = v33
            _ = v34
        }
    }
    switch "foo bar" {
    case "":
        var v35 int8 = 42
        _ = v35
    }
    v10 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
