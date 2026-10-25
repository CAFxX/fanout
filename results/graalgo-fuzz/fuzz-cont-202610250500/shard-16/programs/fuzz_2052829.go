package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 22
    var v3 int64 = -278
    v4 := make(map[string]bool)
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int32 = -2147483648
        _ = fi5
        _ = v6
    }
    v3 += (v3 << ((65535) & 63))
    v3 -= 0
    v7 := v1["\n\t"]
    var v8 uint16 = uint16(v2)
    if true {
        if (42 == v3) {
            switch int(v7) {
            case 2147483647:
                v2 -= 0
            case 32768:
                switch "\n\t" {
                case "\n\t":
                    if true {
                        v9 := v1["hello"]
                        _ = v9
                    } else {
                        var v10 bool = true
                        switch ("hello" + "") {
                        case "foo bar":
                            fmt.Println(v8, v10)
                        case "":
                            v8 = (v8 & v8)
                        case "hello":
                            var v11 uint64 = 824
                            _ = v11
                        default:
                            var v12 string = "hello"
                            _ = v12
                        }
                        switch ((("" + "hello") + ("foo bar" + ("foo bar" + ""))) + "hello") {
                        case "hello":
                            v13 := []int{1}
                            _ = v13
                        case "foo bar":
                            var v14 int8 = -92
                            _ = v14
                        }
                        _ = v10
                    }
                case "foo bar":
                    v15 := GenS1{F0: 4294967295, F1: (false || (uint(v2) > 0)), F2: v3}
                    _ = v15
                case "hello":
                    if true {
                        v4["foo bar"] = (((v8 - v8) > v8) && (-2.25 <= 1e10))
                    }
                }
            default:
                for fi16 := 0; fi16 < 9; fi16++ {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        switch "hello" {
                        case "hello":
                            fmt.Println(v8, v7, fi16)
                        }
                        v7 += v7
                        var v18 string = "hello"
                        _ = fi17
                        _ = v18
                    }
                    v7 -= v7
                    _ = fi16
                }
            }
        }
        switch "a" {
        case "\n\t":
            v8 = (v8 >> ((0) & 63))
        case "hello":
            var v19 int32 = -238
            _ = v19
        case "a":
            _ = len(v4)
        }
        v20 := []int{}
        _ = v20
    } else {
        v2 -= (v2 - 71)
        if false {
            var v21 bool = (v2 >= (v2 + (v2 & 56)))
            v22 := GenS1{F0: 18446744073709551615, F1: (!v21), F2: 1}
            v3 = v3
            _ = v21
            _ = v22
        }
    }
    var v23 uint32 = (^(v7))
    fmt.Println(v3, v3)
    var v24 float64 = -2.25
    v4["foo bar"] = (float64(-2.25) <= v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
