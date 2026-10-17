package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 bool }

func main() {
    var v1 uint8 = 255
    var v2 float32 = 1e10
    var v3 int8 = 42
    v4 := make(map[string]uint32)
    var v5 int64 = 122
    var v6 float64 = (((1.5 / 3.14159) - 0.0) / 1.0)
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 float32 = (v2 * ((v2 / 3.14159) / v2))
        for fi9 := 0; fi9 < 7; fi9++ {
            v8 = v8
            v10 := GenS1{F0: (-1 > -370), F1: 181, F2: true}
            fmt.Println(fi7)
            if true {
                v5 -= -1
            }
            var v11 uint16 = 1
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := v4[""]
        _ = fi7
        _ = v12
        _ = v8
    }
    if (true && true) {
        fmt.Println(v2, v5)
    } else {
        if ((v5 ^ v5) <= v5) {
            fmt.Println(v2, v1, v6)
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := v4["a"]
                _ = fi13
                _ = v14
            }
        }
        var v15 int32 = -2147483648
        _ = v15
    }
    v5 = -645
    switch 897 {
    case 4294967295:
        var v16 uint64 = 4294967295
        _ = v16
    case -569:
        if (!true) {
            var v17 uint16 = 2
            var v18 int32 = -232
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 int32 = int32(v1)
                v4[("hello" + "\n\t")] = 42
                _ = fi19
                _ = v20
            }
            _ = v17
            _ = v18
        } else {
            var v21 uint = 849
            v22 := v4[""]
            _ = v21
            _ = v22
        }
    default:
        v1 -= (v1 * v1)
    }
    if false {
        var v23 int16 = -790
        fmt.Println(v1, v6, v23)
        _ = v23
    } else {
        switch 42 {
        case -2147483648:
            v6 += 1e10
        case 896:
            fmt.Println(v5)
        }
    }
    v24 := GenS1{F0: false, F1: -506, F2: (("\n\t" + "hello") == ("a" + "a"))}
    v3 -= -1
    if (v6 <= 0.0) {
        v25 := []int{295}
        var v26 uint64 = 501
        if ((v3 << ((4294967295) & 63)) <= -1) {
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            fmt.Println(v3, v26)
            for fi27 := 0; fi27 < 10; fi27++ {
                switch "" {
                case "foo bar":
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                case "\n\t":
                    fmt.Println(v2, v1)
                case "hello":
                    if (2 <= 784) {
                        v28 := v4["\n\t"]
                        switch ("hello" + ("" + (("hello" + "a") + "a"))) {
                        case "":
                            fmt.Println(fi27, v28)
                        }
                        _ = v28
                    }
                default:
                    var v29 uint = 42
                    _ = v29
                }
                for fi30 := 0; fi30 < 18; fi30++ {
                    var v31 int32 = -371
                    if (891 >= 4294967295) {
                        fmt.Println(fi30, v2, fi30)
                        if ((0 <= 2) && false) {
                            var v32 float32 = v2
                            _ = v32
                        }
                        var v33 float32 = v2
                        var v34 uint16 = 32767
                        var v35 int8 = v3
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    var v36 bool = (true && false)
                    _ = fi30
                    _ = v31
                    _ = v36
                }
                v37 := GenS1{F0: true, F1: 761, F2: true}
                _ = fi27
                _ = v37
            }
        } else {
            fmt.Println(v26, v26)
            var v38 bool = true
            _ = v38
        }
        _ = v25
        _ = v26
    }
    delete(v4, ("foo bar" + ((("foo bar" + ("foo bar" + ("hello" + ("a" + "a")))) + ("\n\t" + "a")) + "")))
    for fi39 := 0; fi39 < 12; fi39++ {
        var v40 string = "\n\t"
        _ = fi39
        _ = v40
    }
    var v41 float64 = (((v6 * float64((v2 + 0.0))) / 1.5) * 1.5)
    v3 = v3
    if (255 >= 2) {
        v5 = (v5 / (((-(v5))) | 1))
        var v42 int = -777
        _ = v42
    }
    var v43 int32 = 0
    v1 = 246
    if true {
        fmt.Println(v6)
        v44 := GenS1{F0: (v43 != v43), F1: 1000, F2: true}
        _ = v44
    } else {
        fmt.Println(v6)
    }
    fmt.Println(v43, v5, v5)
    if (!(false || (42 <= 367))) {
        fmt.Println(v6, v3, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v5
    _ = v6
}
