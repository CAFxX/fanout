package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }
type GenS2 struct { F0 int; F1 uint16; F2 int16 }
type GenS3 struct { F0 int64; F1 int64 }
func genFn0(p0 uint32) bool {
    var v5 float32 = 3.14159
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS1{F0: 1, F1: -32768, F2: (("\n\t" + "hello") + ((("\n\t" + "hello") + "") + ""))}
            _ = fi7
            _ = v8
        }
        fmt.Println(fi6, p0)
        var v9 uint16 = 1
        _ = fi6
        _ = v9
    }
    _ = p0
    _ = v5
    return true
}

func main() {
    var v1 uint8 = 1
    var v2 int8 = 1
    var v3 string = (("hello" + ((((("hello" + "foo bar") + "\n\t") + "") + "\n\t") + "")) + ("" + "foo bar"))
    var v4 int32 = 0
    v10 := []int{-350}
    for fi11 := 0; fi11 < 19; fi11++ {
        if len(v10) > 0 {
            v10[0] = fi11
        }
        var v12 int16 = 32767
        v13 := GenS1{F0: 42, F1: (^(v12)), F2: ""}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint32 = 374
    switch v3 {
    case "hello":
        fmt.Println(v14, v1, v1)
    case "foo bar":
        if true {
            switch v3 {
            case "":
                var v15 uint8 = (^(v1))
                _ = v15
            default:
                fmt.Println(v14, v1)
            }
        } else {
            for fi16 := 0; fi16 < 12; fi16++ {
                v1 -= (^(v1))
                _ = fi16
            }
        }
    default:
        switch 361 {
        case 127:
            v3 = (v3 + v3)
        default:
            var v17 string = v3
            _ = v17
        }
    }
    fmt.Println(v1, v1, v1)
    var v18 uint32 = 1
    var v19 int16 = -449
    if true {
        var v20 bool = (42 >= (v1 >> ((18446744073709551615) & 63)))
        v4 = 130
        _ = v20
    } else {
        var v21 uint8 = (v1 + (v1 << ((0) & 63)))
        _ = len(v10)
        _ = v21
    }
    var v22 uint16 = 255
    switch 0 {
    case 0:
        var v23 uint64 = 616
        _ = v23
    case -731:
        var v24 string = v3
        _ = v24
    case 127:
        v25 := GenS2{F0: -955, F1: 0, F2: v19}
        _ = v25
    default:
        var v26 uint8 = 75
        _ = v26
    }
    v10 = append(v10, 82)
    switch -2147483648 {
    case 128:
        switch v3 {
        case "foo bar":
            var v27 float64 = -2.25
            _ = v27
        case "a":
            fmt.Println(v2)
        }
    }
    v28 := GenS3{F0: 9223372036854775807, F1: 42}
    if true {
        var v29 uint = 65535
        _ = v29
    } else {
        if false {
            v30 := GenS1{F0: 0, F1: -451, F2: v3}
            for fi31 := 0; fi31 < 19; fi31++ {
                if false {
                    var v32 uint32 = 2147483647
                    if len(v10) > 0 {
                        v10[0] = (fi31 & int(v22))
                    }
                    if true {
                        v1 = 205
                        switch (-(fi31)) {
                        case 4294967295:
                            v4 = 0
                        case -301:
                            for fi33 := 0; fi33 < 12; fi33++ {
                                var v34 int32 = v4
                                _ = fi33
                                _ = v34
                            }
                        default:
                            fmt.Println(v32)
                        }
                        var v35 int64 = 42
                        _ = v35
                    } else {
                        if (true || true) {
                            v3 = "hello"
                        }
                        var v36 uint = 2
                        _ = v36
                    }
                    _ = v32
                }
                _ = fi31
            }
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = v30
        } else {
            _ = len(v10)
            v3 = v3
            for fi37 := 0; fi37 < 4; fi37++ {
                var v38 string = v3
                _ = fi37
                _ = v38
            }
        }
    }
    var v39 int8 = -47
    if ((v18 | (^(v18))) < 650) {
        for fi40 := 0; fi40 < 12; fi40++ {
            v39 = 118
            if (!((v22 ^ 32767) > 958)) {
                if len(v10) > 0 {
                    v10[0] = (fi40 | fi40)
                }
                v41 := GenS2{F0: fi40, F1: 42, F2: -32768}
                var v42 int32 = (v4 % ((v4) | 1))
                _ = v41
                _ = v42
            } else {
                var v43 int16 = (v19 - 1)
                _ = v43
            }
            if true {
                v44 := GenS1{F0: 28, F1: v19, F2: (v3 + ("a" + "a"))}
                _ = v44
            }
            _ = fi40
        }
        v14 = 4294967295
        switch int(v1) {
        case -676:
            switch 2147483647 {
            case 32768:
                v3 = (v3 + "\n\t")
            case -9223372036854775808:
                var v45 int16 = -169
                _ = v45
            default:
                var v46 uint16 = (v22 ^ v22)
                _ = v46
            }
        case 603:
            fmt.Println(v1, v39)
        case 127:
            var v47 int64 = 861
            _ = v47
        default:
            v39 -= (v39 - (^(v39)))
        }
    }
    v10 = append(v10, 905)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v39
    _ = v4
}
