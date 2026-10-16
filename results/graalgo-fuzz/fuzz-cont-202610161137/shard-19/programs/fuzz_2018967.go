package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int64, p1 bool) int64 {
    p0 -= (-(p0))
    var v4 int32 = 762
    for fi5 := 0; fi5 < 1; fi5++ {
        p0 += p0
        var v6 uint = 1
        _ = fi5
        _ = v6
    }
    if (true || false) {
        var v7 string = ("" + (("\n\t" + "foo bar") + "foo bar"))
        var v8 int = -462
        _ = v7
        _ = v8
    } else {
        if true {
            fmt.Println(p0)
            var v9 int32 = v4
            v10 := []int{-434}
            _ = v10
            _ = v9
        }
        v11 := []int{42, 2147483647}
        _ = v11
    }
    if (p0 >= (-(p0))) {
        fmt.Println(v4, p1)
        var v12 bool = ("" < "")
        _ = v12
    } else {
        v13 := []int{-1, -9223372036854775808, 128, 127}
        switch 128 {
        case 127:
            var v14 int16 = 885
            _ = v14
        }
        if len(v13) > 0 {
            v13[0] = 915
        }
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v4
    return p0
}

func main() {
    var v1 float32 = 3.14159
    var v2 string = ("a" + "foo bar")
    v3 := []int{9223372036854775807, -9223372036854775808}
    switch v2 {
    case "\n\t":
        v1 += float32(1.5)
    case "hello":
        fmt.Println(v1, v1)
    default:
        var v15 uint16 = 988
        _ = v15
    }
    var v16 uint = 647
    switch 127 {
    case -9223372036854775808:
        var v17 float64 = (0.1 * 3.14159)
        _ = v17
    case 194:
        v2 = v2
    case -825:
        var v18 uint = 1
        _ = v18
    }
    fmt.Println(v16, v1)
    if (217 < 128) {
        for fi19 := 0; fi19 < 12; fi19++ {
            if (true && (127 != 255)) {
                var v20 int64 = 42
                v20 -= v20
                v21 := GenS1{F0: 106}
                _ = v20
                _ = v21
            } else {
                for fi22 := 0; fi22 < 3; fi22++ {
                    if true {
                        var v23 int16 = 1000
                        switch fi22 {
                        case -366:
                            var v24 int8 = 110
                            _ = v24
                        case 127:
                            var v25 uint = v16
                            switch ("\n\t" + v2) {
                            case "a":
                                switch (v2 + "hello") {
                                case "":
                                    fmt.Println(v25, v16, v16)
                                default:
                                    v2 = v2
                                }
                            case "hello":
                                var v26 uint32 = 2
                                _ = v26
                            case "foo bar":
                                if (!false) {
                                    fmt.Println(fi19, fi22)
                                    var v27 bool = (0 == -2147483648)
                                    v23 = v23
                                    _ = len(v3)
                                    _ = v27
                                } else {
                                    v2 = (v2 + (v2 + v2))
                                }
                            default:
                                fmt.Println(v1, v2)
                            }
                            _ = v25
                        }
                        _ = v23
                    } else {
                        switch fi22 {
                        case -8:
                            var v28 string = v2
                            _ = v28
                        case 1:
                            switch ("a" + "foo bar") {
                            case "":
                                v16 += (v16 | v16)
                            case "hello":
                                _ = len(v3)
                            }
                        }
                        switch "" {
                        case "\n\t":
                            v29 := GenS1{F0: 1}
                            _ = v29
                        case "hello":
                            var v30 int64 = 0
                            _ = v30
                        }
                        var v31 float64 = ((-2.25 / (1.5 / (1e10 - 3.14159))) * 3.14159)
                        _ = v31
                    }
                    _ = fi22
                }
                for fi32 := 0; fi32 < 15; fi32++ {
                    fmt.Println(fi32, v2)
                    var v33 float64 = (0.1 - 0.0)
                    _ = fi32
                    _ = v33
                }
            }
            var v34 uint8 = 255
            _ = fi19
            _ = v34
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if false {
            v35 := GenS1{F0: 45}
            _ = v35
        }
    }
    v36 := GenS1{F0: -1}
    switch -561 {
    case -9223372036854775808:
        fmt.Println(v2, v16)
    case 4294967295:
        var v37 int32 = 42
        _ = v37
    }
    if false {
        for fi38 := 0; fi38 < 14; fi38++ {
            v39 := GenS1{F0: 100}
            v2 = v2
            _ = fi38
            _ = v39
        }
        v3 = append(v3, -9223372036854775808)
    }
    v3 = append(v3, -804)
    v2 = v2
    if true {
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v36
}
