package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int32) bool {
    v5 := []int{}
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int8 = 100
        if false {
            switch "" {
            case "\n\t":
                v8 := GenS1{F0: true}
                _ = v8
            case "a":
                for fi9 := 0; fi9 < 2; fi9++ {
                    var v10 int64 = -9223372036854775808
                    var v11 int16 = 1
                    if false {
                        switch (fi9 & (-(fi6))) {
                        case 127:
                            _ = len(v5)
                        case -9223372036854775808:
                            var v12 int16 = -32768
                            v11 += v12
                            _ = v12
                        case 0:
                            var v13 uint32 = 42
                            _ = v13
                        default:
                            var v14 int = (-(fi9))
                            _ = v14
                        }
                    } else {
                        switch "hello" {
                        case "a":
                            if (272 == 52) {
                                v15 := GenS1{F0: true}
                                var v16 int = -1
                                var v17 bool = ((!(128 <= 255)) && false)
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                        }
                        v7 += 42
                    }
                    v18 := GenS1{F0: (false || true)}
                    v19 := GenS1{F0: (!(fi9 != (fi9 % ((fi9) | 1))))}
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v18
                    _ = v19
                }
            case "hello":
                var v20 int64 = 34
                _ = v20
            }
            for fi21 := 0; fi21 < 3; fi21++ {
                var v22 float32 = ((-2.25 + 1.5) - 3.14159)
                _ = fi21
                _ = v22
            }
            var v23 uint64 = 4294967295
            _ = v23
        } else {
            var v24 int64 = 519
            _ = v24
        }
        _ = fi6
        _ = v7
    }
    _ = len(v5)
    switch ("foo bar" + "") {
    case "foo bar":
        fmt.Println(p0)
    }
    _ = p0
    _ = v5
    return false
}
func genFn1() uint64 {
    if true {
        var v25 int8 = 60
        v26 := []int{32768}
        _ = len(v26)
        _ = v25
        _ = v26
    }
    var v27 int32 = 2147483647
    var v28 int64 = 0
    var v29 int32 = (v27 + -2147483648)
    _ = v27
    _ = v28
    _ = v29
    return 42
}

func main() {
    var v1 int8 = 42
    var v2 uint32 = 4294967295
    var v3 uint64 = 1
    v4 := make(map[string]string)
    v30 := GenS1{F0: true}
    var v31 int64 = 42
    v32 := []int{}
    switch "" {
    case "foo bar":
        v33 := GenS1{F0: (!false)}
        _ = v33
    }
    var v34 int8 = v1
    for fi35 := 0; fi35 < 10; fi35++ {
        var v36 uint8 = 128
        v37 := GenS1{F0: (!false)}
        var v38 uint16 = 434
        _ = fi35
        _ = v36
        _ = v37
        _ = v38
    }
    var v39 string = ("\n\t" + "foo bar")
    var v40 string = ("\n\t" + ((v39 + (v39 + (v39 + ((v39 + "\n\t") + ((v39 + v39) + (v39 + v39)))))) + v39))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v39
    _ = v4
    _ = v40
}
