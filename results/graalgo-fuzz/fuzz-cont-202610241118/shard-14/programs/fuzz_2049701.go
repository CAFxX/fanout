package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 bool, p1 int64) uint32 {
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint8 = 42
        var v9 int8 = 32
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 724
    p0 = p0
    _ = p0
    _ = p1
    _ = v10
    return v10
}

func main() {
    var v1 uint16 = 1
    v2 := []int{470, 843, 894}
    v3 := make(map[string]int)
    var v4 uint32 = 4294967295
    v5 := make(map[int]int)
    var v6 bool = false
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int = (^(fi11))
        v12 -= -2147483648
        v6 = v6
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: -1}
    v14 := v5[521]
    var v15 uint16 = (^(v1))
    fmt.Println(v15, v15, v1)
    delete(v5, v14)
    var v16 string = ("hello" + "")
    for fi17 := 0; fi17 < 12; fi17++ {
        if (((v16 + "\n\t") + (("hello" + ("" + v16)) + (v16 + v16))) <= ((v16 + v16) + (v16 + ((v16 + v16) + (("foo bar" + v16) + v16))))) {
            v14 -= -9223372036854775808
            fmt.Println(v6, v4)
        } else {
            if false {
                for fi18 := 0; fi18 < 7; fi18++ {
                    v2 = append(v2, (v14 | fi18))
                    v14 -= fi18
                    var v19 bool = ((3.14159 + -2.25) < (3.14159 + 0.1))
                    _ = fi18
                    _ = v19
                }
            } else {
                var v20 float64 = 3.14159
                _ = v20
            }
            var v21 uint = 65535
            var v22 int64 = -1
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 13; fi23++ {
            v1 = (v15 & v15)
            if v6 {
                if (v6 && true) {
                    var v24 uint8 = 233
                    v4 = v4
                    _ = v24
                }
                var v25 uint64 = uint64(v4)
                switch v16 {
                case "foo bar":
                    fmt.Println(fi23, v4)
                case "":
                    v26 := GenS2{F0: 2147483648}
                    _ = v26
                }
                _ = v25
            }
            _ = fi23
        }
        _ = fi17
    }
    var v27 int8 = 89
    if v6 {
        v15 = (v1 & v15)
        var v28 int64 = int64(v14)
        v27 = v27
        _ = v28
    } else {
        switch v16 {
        case "hello":
            if (v6 && v6) {
                if (!(((v4 == v4) && false) && true)) {
                    fmt.Println(v14, v27, v4)
                    var v29 uint8 = 128
                    v16 = "hello"
                    _ = v29
                } else {
                    _ = len(v2)
                    var v30 int = v14
                    _ = v30
                }
                var v31 float64 = (3.14159 - (-2.25 + 0.1))
                v14 -= 9223372036854775807
                _ = v31
            } else {
                v32 := GenS2{F0: v4}
                _ = v32
            }
        default:
            switch v16 {
            case "hello":
                switch -11 {
                case -2147483648:
                    v15 = (v1 & v1)
                case 4294967295:
                    var v33 string = v16
                    _ = v33
                case -883:
                    v27 = (v27 | v27)
                default:
                    v16 = v16
                }
            case "foo bar":
                v15 += v1
            case "a":
                switch (v16 + "hello") {
                case "hello":
                    v27 = v27
                case "a":
                    var v34 uint64 = 474
                    _ = v34
                case "\n\t":
                    var v35 float64 = ((3.14159 - 1.5) + (3.14159 + 0.1))
                    _ = v35
                }
            default:
                if true {
                    for fi36 := 0; fi36 < 1; fi36++ {
                        _ = len(v3)
                        v16 = ((v16 + ((("foo bar" + v16) + v16) + ("foo bar" + ((v16 + (v16 + (v16 + (v16 + v16)))) + "foo bar")))) + v16)
                        var v37 string = (v16 + ((v16 + v16) + (v16 + ("hello" + v16))))
                        _ = fi36
                        _ = v37
                    }
                    v38 := v3[""]
                    _ = v38
                }
            }
        }
        switch 42 {
        case 2147483647:
            var v39 int16 = 431
            _ = v39
        }
    }
    v4 += (v4 * v4)
    var v40 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
