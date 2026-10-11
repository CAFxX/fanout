package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }

func main() {
    v1 := []uint32{2147483648, 1, 196}
    var v2 int64 = 42
    var v3 int64 = v2
    var v4 int8 = 104
    var v5 uint = 18446744073709551615
    v6 := make(map[int]int)
    v7 := GenS1{F0: ""}
    fmt.Println(v4, v2)
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 string = "a"
        if true {
            fmt.Println(v3, v4)
            switch "\n\t" {
            case "foo bar":
                if false {
                    var v10 float64 = 0.1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v10
                }
            case "a":
                v1 = append(v1, 4294967295)
            }
            var v11 float32 = 3.14159
            _ = v11
        } else {
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(v5)
                if false {
                    switch "" {
                    case "hello":
                        v13 := GenS1{F0: ""}
                        _ = v13
                    case "foo bar":
                        var v14 uint32 = 42
                        _ = v14
                    case "a":
                        v15 := GenS2{F0: (v3 % ((v3) | 1))}
                        _ = v15
                    }
                }
                _ = fi12
            }
            var v16 int32 = int32(fi8)
            _ = v16
        }
        _ = fi8
        _ = v9
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 13; fi18++ {
            v19 := GenS1{F0: "foo bar"}
            fmt.Println(fi18, v4, fi17)
            if false {
                _ = len(v6)
            }
            switch fi17 {
            case 243:
                var v20 uint32 = 1
                _ = v20
            }
            if false {
                fmt.Println(v5)
                switch "\n\t" {
                case "\n\t":
                    var v21 int64 = -862
                    _ = v21
                }
                var v22 float64 = 1e10
                _ = v22
            }
            fmt.Println(v3)
            _ = fi18
            _ = v19
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 int64 = -9223372036854775808
            v25 := GenS1{F0: "hello"}
            _ = fi23
            _ = v24
            _ = v25
        }
        if true {
            v26 := GenS3{F0: 127}
            if (42 < v4) {
                var v27 bool = ("" >= "hello")
                v28 := GenS2{F0: 1}
                switch fi17 {
                case 128:
                    for fi29 := 0; fi29 < 13; fi29++ {
                        var v30 int = 2147483647
                        if (0.1 < -2.25) {
                            var v31 uint = (v5 / ((891) | 1))
                            v30 += fi29
                            if len(v1) > 0 {
                                v1[0] = 2147483647
                            }
                            _ = v31
                        }
                        _ = fi29
                        _ = v30
                    }
                case 393:
                    if false {
                        delete(v6, fi17)
                        for fi32 := 0; fi32 < 4; fi32++ {
                            v4 += (v4 / ((v4) | 1))
                            _ = fi32
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        v33 := GenS2{F0: (v2 + (v3 & 0))}
                        v34 := GenS1{F0: (((("hello" + "a") + "foo bar") + "hello") + (("foo bar" + "\n\t") + ""))}
                        fmt.Println(v3)
                        _ = v33
                        _ = v34
                    }
                default:
                    var v35 bool = v27
                    _ = v35
                }
                _ = v27
                _ = v28
            }
            if true {
                fmt.Println(v3)
                var v36 bool = (65535 <= 32767)
                var v37 float64 = 1.5
                _ = v36
                _ = v37
            } else {
                fmt.Println(v3, v5)
            }
            _ = v26
        }
        _ = fi17
    }
    if (!false) {
        if true {
            var v38 uint32 = 2
            for fi39 := 0; fi39 < 5; fi39++ {
                var v40 int = fi39
                _ = fi39
                _ = v40
            }
            _ = v38
        } else {
            fmt.Println(v3)
        }
        v5 = 255
        var v41 uint64 = 151
        _ = v41
    }
    var v42 uint8 = 127
    v42 = (^(v42))
    var v43 uint32 = 199
    v44 := GenS1{F0: "hello"}
    if len(v1) > 0 {
        v1[0] = 68
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
