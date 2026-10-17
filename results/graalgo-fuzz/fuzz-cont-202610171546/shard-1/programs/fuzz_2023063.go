package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 bool; F2 int64 }
func genFn0(p0 uint8) uint64 {
    var v6 float64 = (-2.25 + ((float64(1.5) + 0.0) + float64(0.1)))
    switch -297 {
    case 2147483647:
        var v7 int32 = -612
        _ = v7
    case -153:
        switch 790 {
        case 2147483647:
            v8 := GenS1{F0: 956}
            _ = v8
        case -337:
            v6 += ((0.1 - v6) + v6)
        case -152:
            switch ("hello" + (("\n\t" + "") + ((("foo bar" + ("\n\t" + "foo bar")) + ("" + "")) + "foo bar"))) {
            case "\n\t":
                if ((p0 + 1) != p0) {
                    v9 := []int{-555, 440, -368}
                    var v10 int16 = 1
                    _ = v10
                    _ = v9
                }
            case "hello":
                var v11 uint = 255
                _ = v11
            }
        }
    }
    switch "hello" {
    case "":
        v12 := []int{9223372036854775807}
        _ = v12
    case "hello":
        v13 := GenS1{F0: 18446744073709551615}
        _ = v13
    default:
        for fi14 := 0; fi14 < 12; fi14++ {
            v6 -= 0.1
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 uint16 = 65535
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    v6 = v6
    v17 := []int{}
    _ = p0
    _ = v17
    _ = v6
    return 363
}
func genFn1(p0 uint8, p1 int16, p2 uint32) bool {
    v18 := []int{-9223372036854775808}
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(p0, p2, p2)
        _ = fi19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return (!false)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 42
    var v3 bool = (!(46 < 42))
    v4 := []string{"hello", ("hello" + "hello"), ((("foo bar" + ("foo bar" + ("hello" + "hello"))) + "a") + ("\n\t" + "\n\t")), "a"}
    var v5 uint16 = 1
    var v20 bool = (!(!(-753 == 1)))
    switch 9223372036854775807 {
    case -500:
        v21 := GenS1{F0: 4294967295}
        _ = v21
    case -9223372036854775808:
        var v22 uint32 = 0
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 uint16 = (v5 / ((65535) | 1))
        if true {
            if v3 {
                fmt.Println(v3, v20, v3)
                if v3 {
                    v5 -= 255
                }
                v20 = (v20 && (v20 && false))
            } else {
                if (!v20) {
                    v25 := GenS1{F0: 2}
                    if true {
                        var v26 uint64 = 2
                        var v27 uint = 891
                        var v28 int16 = v2
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 9; fi29++ {
                        switch (fi29 / ((int(v5)) | 1)) {
                        case 4294967295:
                            fmt.Println(fi23)
                        }
                        var v30 float64 = 0.0
                        v20 = v20
                        _ = fi29
                        _ = v30
                    }
                    _ = v25
                }
            }
            var v31 float32 = 1.5
            for fi32 := 0; fi32 < 15; fi32++ {
                var v33 string = "a"
                _ = fi32
                _ = v33
            }
            _ = v31
        }
        _ = fi23
        _ = v24
    }
    var v34 float32 = 0.0
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    fmt.Println(v34, v5, v34)
    var v35 int = -939
    v36 := GenS2{F0: 0, F1: false, F2: -424}
    var v37 bool = v20
    var v38 int8 = 100
    v39 := GenS2{F0: v2, F1: v3, F2: 9223372036854775807}
    var v40 int = -2147483648
    v37 = (!(!((-863 >= v2) && (v20 && (!v3)))))
    var v41 int64 = int64(v2)
    switch v35 {
    case -820:
        var v42 uint8 = 2
        _ = v42
    case 42:
        v5 -= 42
    default:
        var v43 float64 = 3.14159
        _ = v43
    }
    v44 := GenS1{F0: 65535}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v40, v2, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v5
}
