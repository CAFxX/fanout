package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int; F2 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    var v3 string = ("a" + "foo bar")
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    v3 = v3
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint = 2
            var v9 int16 = -609
            var v10 float64 = 1.5
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            if false {
                v3 = "\n\t"
                var v12 uint = 68
                v3 = v3
                _ = v12
            } else {
                var v13 int8 = 0
                if false {
                    v13 = (v13 >> ((579) & 63))
                    var v14 bool = false
                    _ = v14
                }
                _ = v13
            }
            var v15 int8 = 42
            _ = fi11
            _ = v15
        }
        _ = fi6
    }
    v16 := GenS1{F0: (v3 + v3), F1: 4294967295}
    if true {
        var v17 uint16 = 0
        _ = v17
    } else {
        for fi18 := 0; fi18 < 5; fi18++ {
            fmt.Println(fi18, fi18)
            _ = fi18
        }
    }
    fmt.Println(v3)
    v3 = (((v3 + "a") + v3) + v3)
    switch v3 {
    case "\n\t":
        v1[(("a" + (v3 + "a")) + (("\n\t" + v3) + v3))] = 767
    case "a":
        _ = len(v4)
    }
    if true {
        v19 := GenS1{F0: "\n\t", F1: 42}
        _ = v19
    }
    if (!(!false)) {
        v3 = v3
        switch -9223372036854775808 {
        case 9223372036854775807:
            var v20 uint8 = 42
            _ = v20
        case 42:
            if ((2 <= 65535) && true) {
                if (255 >= 999) {
                    v3 = "a"
                } else {
                    fmt.Println(v3)
                    var v21 int16 = -955
                    var v22 uint8 = 255
                    _ = v21
                    _ = v22
                }
            } else {
                for fi23 := 0; fi23 < 20; fi23++ {
                    var v24 int8 = 127
                    var v25 string = v3
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                v4[32768] = 32768
            }
        case 396:
            v3 = (v3 + (("" + "a") + v3))
        }
        if (!true) {
            v3 = "hello"
            fmt.Println(v3)
            v26 := []int{4294967295}
            _ = v26
        } else {
            var v27 uint8 = 2
            switch "foo bar" {
            case "\n\t":
                for fi28 := 0; fi28 < 20; fi28++ {
                    var v29 uint32 = 0
                    _ = fi28
                    _ = v29
                }
            default:
                if true {
                    var v30 uint8 = 255
                    for fi31 := 0; fi31 < 10; fi31++ {
                        v27 += v30
                        switch v3 {
                        case "hello":
                            v30 += v30
                        case "\n\t":
                            fmt.Println(fi31)
                        default:
                            var v32 uint16 = 686
                            _ = v32
                        }
                        v33 := GenS2{F0: 585}
                        _ = fi31
                        _ = v33
                    }
                    _ = v30
                } else {
                    _ = len(v4)
                }
            }
            if ((false || ((v27 >> ((0) & 63)) >= (v27 / ((121) | 1)))) && false) {
                var v34 int32 = int32(v27)
                var v35 int8 = -86
                _ = v34
                _ = v35
            } else {
                if ((((v3 + "foo bar") + "") + "a") != (("" + "") + v3)) {
                    var v36 uint8 = 255
                    _ = v36
                }
                v27 += v27
                var v37 string = v3
                _ = v37
            }
            _ = v27
        }
    }
    v38 := GenS3{F0: -1, F1: 32767, F2: 0}
    if true {
        v3 = v3
    } else {
        switch v3 {
        case "":
            var v39 float64 = 1.5
            _ = v39
        default:
            var v40 uint32 = 2147483648
            _ = v40
        }
        switch 2147483647 {
        case -1:
            v3 = ("foo bar" + "a")
        default:
            v41 := []int{2147483647, -541, 42, 127}
            _ = v41
        }
    }
    fmt.Println(v3)
    v42 := []int{-9223372036854775808}
    var v43 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
