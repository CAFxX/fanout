package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int8; F2 int32 }
type GenS3 struct { F0 bool; F1 int; F2 string }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]string)
    var v3 bool = (("foo bar" + "foo bar") < ("\n\t" + "foo bar"))
    if true {
        fmt.Println(v3, v1)
        v1 = (v1 % (((v1 | v1)) | 1))
        v4 := []int{}
        _ = v4
    }
    v1 = (v1 & 0)
    if (-760 > 2147483647) {
        v5 := GenS1{F0: 106}
        var v6 float32 = 1e10
        v7 := []int{2147483647, 2147483647}
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 bool = (0 >= -897)
    v9 := GenS1{F0: 832}
    var v10 uint8 = v1
    var v11 int64 = 0
    switch -2147483648 {
    case -9223372036854775808:
        for fi12 := 0; fi12 < 14; fi12++ {
            v8 = false
            v13 := GenS2{F0: 42, F1: -128, F2: 1}
            _ = fi12
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 18; fi14++ {
            if (v8 || (1 > 227)) {
                v15 := GenS1{F0: 42}
                v16 := []int{fi14, (fi14 + fi14), fi14}
                _ = v15
                _ = v16
            } else {
                if v8 {
                    if (100 == -128) {
                        v3 = v8
                        v17 := GenS2{F0: fi14, F1: 127, F2: 597}
                        _ = v17
                    }
                } else {
                    v1 -= (^(v1))
                }
                v2["a"] = "foo bar"
            }
            v18 := []int{}
            _ = fi14
            _ = v18
        }
    }
    v3 = (-2147483648 != 32768)
    switch 127 {
    case 59:
        if (2 == 65535) {
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := GenS1{F0: 1}
                _ = fi19
                _ = v20
            }
            if v8 {
                v21 := []int{}
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
                v22 := GenS1{F0: 1}
                _ = v21
                _ = v22
            }
        } else {
            for fi23 := 0; fi23 < 18; fi23++ {
                fmt.Println(fi23, v1)
                switch (("" + "hello") + "foo bar") {
                case "\n\t":
                    delete(v2, "\n\t")
                case "foo bar":
                    var v24 int64 = 295
                    _ = v24
                case "hello":
                    v25 := GenS3{F0: (v1 != v10), F1: fi23, F2: ""}
                    _ = v25
                }
                _ = fi23
            }
            switch 42 {
            case 9223372036854775807:
                v11 = v11
            case 128:
                var v26 int16 = -1
                _ = v26
            case -1:
                var v27 int32 = 2147483647
                _ = v27
            default:
                v8 = v8
            }
            v28 := []int{-2147483648, 1}
            _ = v28
        }
    default:
        for fi29 := 0; fi29 < 18; fi29++ {
            var v30 uint = 42
            switch fi29 {
            case -9223372036854775808:
                var v31 int16 = 42
                _ = v31
            case 32767:
                if v8 {
                    if (v8 || (713 > 42)) {
                        v10 -= 127
                        v32 := GenS1{F0: 18446744073709551615}
                        if (2147483647 <= 172) {
                            v8 = ("\n\t" > "foo bar")
                        }
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 13; fi33++ {
                        v34 := []int{}
                        _ = fi33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 19; fi35++ {
                        v30 -= v30
                        v36 := []int{}
                        v30 += 1
                        _ = fi35
                        _ = v36
                    }
                } else {
                    var v37 int16 = 1
                    _ = v37
                }
            case -433:
                if (v8 && (32767 > 42)) {
                    v1 = (v10 / (((v1 << ((65535) & 63))) | 1))
                    for fi38 := 0; fi38 < 5; fi38++ {
                        v8 = ((422 < 65535) && (-124 == 0))
                        _ = fi38
                    }
                    var v39 int = -9223372036854775808
                    _ = v39
                }
            default:
                _ = len(v2)
            }
            _ = fi29
            _ = v30
        }
    }
    v3 = v3
    v40 := GenS3{F0: v3, F1: 838, F2: ""}
    for fi41 := 0; fi41 < 2; fi41++ {
        v42 := GenS4{F0: 2147483648}
        _ = fi41
        _ = v42
    }
    switch 79 {
    case -986:
        delete(v2, ("hello" + ""))
    case 127:
        var v43 bool = true
        _ = v43
    }
    v10 -= v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v40
    _ = v8
    _ = v9
}
