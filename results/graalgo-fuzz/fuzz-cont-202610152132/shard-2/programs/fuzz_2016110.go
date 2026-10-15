package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 int16; F2 string }

func main() {
    var v1 string = "a"
    var v2 string = "\n\t"
    var v3 bool = true
    var v4 string = "a"
    v5 := make(map[string]string)
    var v6 int8 = 42
    if (v3 || (9223372036854775807 > -718)) {
        v7 := []int{128, 780, 1, -71}
        if false {
            for fi8 := 0; fi8 < 16; fi8++ {
                switch ("" + ("foo bar" + "foo bar")) {
                case "hello":
                    v9 := GenS1{F0: (fi8 + (fi8 / ((fi8) | 1))), F1: 18446744073709551615}
                    _ = v9
                case "a":
                    for fi10 := 0; fi10 < 8; fi10++ {
                        var v11 uint64 = 18446744073709551615
                        v3 = v3
                        if true {
                            _ = len(v7)
                        } else {
                            var v12 float32 = 1.5
                            fmt.Println(v3, v4)
                            _ = v12
                        }
                        _ = fi10
                        _ = v11
                    }
                }
                _ = fi8
            }
        }
        _ = v7
    } else {
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS1{F0: (fi13 | (fi13 - (-(fi13)))), F1: 65535}
            var v15 string = (v1 + ("a" + "foo bar"))
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    fmt.Println(v1, v1)
    for fi16 := 0; fi16 < 3; fi16++ {
        switch fi16 {
        case -2147483648:
            v17 := GenS1{F0: fi16, F1: 65535}
            _ = v17
        case -9223372036854775808:
            v3 = false
        }
        _ = fi16
    }
    if (v3 || (792 >= 1)) {
        v2 = v1
    } else {
        for fi18 := 0; fi18 < 7; fi18++ {
            v6 -= -98
            _ = fi18
        }
        var v19 float64 = 0.0
        _ = v19
    }
    if ((!v3) && v3) {
        var v20 uint = 310
        v21 := GenS1{F0: -119, F1: (v20 % (((v20 | 4294967295)) | 1))}
        v22 := GenS1{F0: -215, F1: v20}
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := v5[v2]
    var v24 uint16 = 546
    v3 = true
    v25 := GenS1{F0: 32767, F1: 42}
    _ = len(v5)
    if (v2 >= v23) {
        switch 34 {
        case 0:
            v26 := []int{-969, 973}
            _ = v26
        case 4294967295:
            _ = len(v5)
        case 2147483647:
            v1 = v1
        }
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        switch (v2 + v4) {
        case "foo bar":
            switch ((((v1 + (v4 + v23)) + ((("" + (("hello" + ("hello" + v1)) + ("foo bar" + v2))) + ("a" + v2)) + ("foo bar" + ((v1 + v4) + v4)))) + v2) + v2) {
            case "":
                v24 += (v24 % (((v24 / ((142) | 1))) | 1))
            case "foo bar":
                v24 += v24
            case "hello":
                fmt.Println(v6)
            }
        }
        if false {
            var v28 uint = 2
            v29 := GenS2{F0: 609}
            for fi30 := 0; fi30 < 19; fi30++ {
                v31 := GenS2{F0: 2147483647}
                fmt.Println(v2, v23, v3)
                var v32 uint64 = 4294967295
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = v28
            _ = v29
        } else {
            switch fi27 {
            case 128:
                var v33 int64 = 0
                _ = v33
            case -2147483648:
                var v34 uint8 = 0
                _ = v34
            case 9223372036854775807:
                switch 915 {
                case -70:
                    v35 := GenS3{F0: (!((v24 >> ((18446744073709551615) & 63)) <= v24)), F1: 1000, F2: ("a" + "")}
                    _ = v35
                default:
                    fmt.Println(v6, v23, v23)
                }
            }
            v2 = ("foo bar" + ("foo bar" + ("" + v2)))
            _ = len(v5)
        }
        _ = fi27
    }
    var v36 uint = 4294967295
    var v37 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
