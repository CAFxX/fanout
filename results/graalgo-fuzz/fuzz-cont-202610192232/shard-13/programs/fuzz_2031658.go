package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint8; F2 int }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 255
    v3 := make(map[string]int)
    var v4 uint = (^(v2))
    v5 := make(map[string]string)
    var v6 int64 = 535
    v7 := GenS1{F0: -1}
    if true {
        v2 = v4
        var v8 bool = (177 >= 42)
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 uint16 = 0
            v11 := []int{(fi9 | (fi9 << ((v2) & 63))), fi9}
            if (v10 < (v10 % (((^(v10))) | 1))) {
                if v8 {
                    switch "foo bar" {
                    case "foo bar":
                        _ = len(v11)
                    case "a":
                        if v8 {
                            if v8 {
                                var v12 int32 = 42
                                delete(v5, "")
                                v6 = 241
                                _ = v12
                            }
                            delete(v1, (fi9 ^ fi9))
                        }
                    case "\n\t":
                        v13 := GenS1{F0: (-(fi9))}
                        _ = v13
                    default:
                        if v8 {
                            fmt.Println(fi9, v10)
                            var v14 int64 = (^(v6))
                            _ = v14
                        } else {
                            switch (fi9 + fi9) {
                            case 2147483647:
                                for fi15 := 0; fi15 < 2; fi15++ {
                                    v2 += 1
                                    var v16 uint16 = v10
                                    _ = fi15
                                    _ = v16
                                }
                            case 32768:
                                switch (("foo bar" + "\n\t") + "foo bar") {
                                case "a":
                                    v10 -= v10
                                case "foo bar":
                                    switch "\n\t" {
                                    case "\n\t":
                                        v17 := GenS2{F0: -2147483648, F1: 172, F2: (fi9 >> ((1) & 63))}
                                        _ = v17
                                    case "hello":
                                        v11 = append(v11, fi9)
                                    }
                                case "":
                                    for fi18 := 0; fi18 < 19; fi18++ {
                                        fmt.Println(v4)
                                        v10 = 42
                                        v19 := GenS1{F0: (^(fi9))}
                                        _ = fi18
                                        _ = v19
                                    }
                                default:
                                    var v20 uint = v4
                                    _ = v20
                                }
                            case -1:
                                v2 = v2
                            default:
                                var v21 bool = v8
                                _ = v21
                            }
                        }
                    }
                } else {
                    var v22 int = 2147483647
                    _ = v22
                }
                var v23 int16 = int16(fi9)
                _ = v23
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    }
    if false {
        var v24 int8 = 0
        fmt.Println(v2)
        v25 := v3[(("" + "") + "hello")]
        _ = v24
        _ = v25
    }
    v26 := GenS1{F0: 9223372036854775807}
    for fi27 := 0; fi27 < 7; fi27++ {
        v2 += v2
        v28 := []int{fi27, -958, 128}
        fmt.Println(v2, v4, fi27)
        _ = fi27
        _ = v28
    }
    delete(v5, ("a" + ((((("\n\t" + "a") + "hello") + "a") + "") + "")))
    if (((!(!true)) || (214 <= 42)) || false) {
        if (!(!false)) {
            for fi29 := 0; fi29 < 8; fi29++ {
                v3[("" + "\n\t")] = fi29
                var v30 uint64 = 18446744073709551615
                if true {
                    var v31 uint64 = (v30 << ((v4) & 63))
                    var v32 int8 = -120
                    _ = v31
                    _ = v32
                } else {
                    fmt.Println(v2)
                    var v33 int8 = -1
                    fmt.Println(v33, v2)
                    _ = v33
                }
                _ = fi29
                _ = v30
            }
            v6 += v6
        }
    }
    switch "\n\t" {
    case "":
        v34 := GenS2{F0: 32767, F1: 184, F2: int(v2)}
        _ = v34
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        switch 127 {
        case 9223372036854775807:
            v36 := []int{fi35, -9223372036854775808, fi35, (fi35 & fi35)}
            _ = v36
        case 4294967295:
            var v37 float32 = 1e10
            _ = v37
        }
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
