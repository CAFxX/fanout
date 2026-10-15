package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 string, p1 uint32) uint {
    var v4 float64 = 1.5
    switch -503 {
    case 0:
        if true {
            var v5 int8 = 42
            p1 += p1
            if (!true) {
                for fi6 := 0; fi6 < 15; fi6++ {
                    switch (("a" + p0) + (p0 + p0)) {
                    case "":
                        for fi7 := 0; fi7 < 14; fi7++ {
                            var v8 int32 = 0
                            fmt.Println(fi7, p1)
                            fmt.Println(v5)
                            var v9 int8 = 127
                            _ = fi7
                            _ = v8
                            _ = v9
                        }
                    case "foo bar":
                        v4 -= 1e10
                    case "a":
                        v10 := []int{(fi6 + fi6), fi6, fi6, fi6}
                        _ = v10
                    }
                    var v11 uint64 = 4294967295
                    _ = fi6
                    _ = v11
                }
                if false {
                    v12 := []int{-834, 2147483647, -287, 32767}
                    var v13 float64 = -2.25
                    v14 := GenS1{F0: 0}
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    fmt.Println(p0)
                    switch 128 {
                    case 1:
                        v15 := []int{}
                        _ = v15
                    case 32768:
                        v16 := []int{-794}
                        _ = v16
                    }
                    v4 = 1.5
                }
                p0 = (((("hello" + (p0 + p0)) + p0) + ((p0 + p0) + p0)) + p0)
            } else {
                if true {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        var v18 int64 = 9223372036854775807
                        fmt.Println(p0, v4)
                        for fi19 := 0; fi19 < 1; fi19++ {
                            var v20 string = ("a" + p0)
                            var v21 uint = 0
                            var v22 int = fi19
                            v23 := []int{}
                            v24 := GenS1{F0: (p1 >> ((v21) & 63))}
                            _ = fi19
                            _ = v20
                            _ = v21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        _ = fi17
                        _ = v18
                    }
                    p1 = p1
                    v25 := GenS1{F0: (p1 + p1)}
                    _ = v25
                }
            }
            _ = v5
        }
    case -734:
        var v26 int8 = 127
        _ = v26
    case -1:
        var v27 int16 = 1
        _ = v27
    default:
        p0 = "a"
    }
    fmt.Println(p0)
    var v28 uint32 = p1
    v4 = v4
    _ = p0
    _ = p1
    _ = v28
    _ = v4
    return 774
}

func main() {
    var v1 int16 = -32768
    var v2 uint = 65535
    var v3 bool = false
    for fi29 := 0; fi29 < 7; fi29++ {
        v30 := []int{}
        fmt.Println(v1)
        _ = fi29
        _ = v30
    }
    var v31 uint = 255
    var v32 uint16 = 0
    fmt.Println(v1)
    var v33 uint8 = 255
    var v34 int32 = 2147483647
    var v35 uint32 = 658
    v36 := GenS1{F0: v35}
    var v37 uint16 = v32
    v34 -= (v34 << ((v31) & 63))
    if v3 {
        switch "hello" {
        case "hello":
            for fi38 := 0; fi38 < 19; fi38++ {
                fmt.Println(v31)
                _ = fi38
            }
        case "foo bar":
            v39 := []int{690, int(v35), -2147483648, -404}
            _ = v39
        default:
            v40 := []int{595, 2147483647, -987, 9223372036854775807}
            _ = v40
        }
        if false {
            var v41 float32 = (1.5 - 1e10)
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        var v43 int16 = (v1 / ((int16(v34)) | 1))
        v44 := GenS1{F0: uint32(v2)}
        _ = fi42
        _ = v43
        _ = v44
    }
    if true {
        v45 := []int{-9223372036854775808, 4294967295}
        _ = v45
    } else {
        if ((^(v35)) < (v35 | (v35 << ((v2) & 63)))) {
            v46 := []int{4294967295, 9223372036854775807}
            _ = v46
        }
        for fi47 := 0; fi47 < 6; fi47++ {
            var v48 float32 = 3.14159
            v37 = (v37 / (((v32 * (v32 + v32))) | 1))
            var v49 uint16 = 255
            _ = fi47
            _ = v48
            _ = v49
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
