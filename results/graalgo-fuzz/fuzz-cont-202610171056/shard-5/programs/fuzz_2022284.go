package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint64; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 uint16; F2 int64 }

func main() {
    var v1 int8 = 1
    v2 := []string{"hello", ("\n\t" + "a"), (("hello" + (("a" + ("a" + "")) + "\n\t")) + "\n\t")}
    var v3 uint64 = 0
    var v4 float32 = float32((0.0 / (3.14159 * 1e10)))
    var v5 uint = 255
    var v6 uint16 = 65535
    var v7 uint = uint(v6)
    if true {
        if ((false || (!(2147483648 <= 65535))) && true) {
            var v8 uint64 = 648
            var v9 string = ("hello" + ("hello" + ("foo bar" + "\n\t")))
            _ = v8
            _ = v9
        }
        v1 = v1
    } else {
        if ((v1 >= v1) || (193 <= 42)) {
            switch (("foo bar" + ("" + "a")) + "") {
            case "hello":
                v7 += (v7 * (v5 | 1))
            case "":
                v10 := GenS1{F0: (1 > 421), F1: 2147483647, F2: 2}
                _ = v10
            case "\n\t":
                switch 356 {
                case -727:
                    v7 = 42
                }
            }
            var v11 uint16 = 65535
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v11
        }
    }
    fmt.Println(v3)
    v12 := GenS2{F0: 1, F1: 0, F2: v1}
    var v13 string = ""
    v2 = append(v2, (v13 + ("a" + (v13 + v13))))
    _ = len(v2)
    var v14 int16 = -32768
    var v15 string = "a"
    fmt.Println(v1, v5, v15)
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            switch v13 {
            case "hello":
                var v18 int32 = -674
                _ = v18
            default:
                switch 2147483647 {
                case 79:
                    v19 := GenS1{F0: (0 >= v14), F1: 1, F2: 230}
                    _ = v19
                case 32768:
                    v6 = v6
                }
            }
            var v20 bool = (((3.14159 / ((3.14159 * ((float64(((v4 * float32(0.1)) + v4)) / 1e10) + float64(0.1))) - 1e10)) * 3.14159) <= 0.0)
            _ = fi17
            _ = v20
        }
        fmt.Println(v1, v14)
        _ = fi16
    }
    v6 = 32767
    for fi21 := 0; fi21 < 20; fi21++ {
        if (!((!false) || false)) {
            if true {
                if false {
                    v13 = v15
                } else {
                    switch (fi21 & (fi21 >> (((^(v7))) & 63))) {
                    case 2147483647:
                        for fi22 := 0; fi22 < 5; fi22++ {
                            v23 := GenS1{F0: (-709 <= -1), F1: 541, F2: 128}
                            _ = fi22
                            _ = v23
                        }
                    case -144:
                        v24 := GenS3{F0: ("\n\t" + ((v13 + ("" + v13)) + (v13 + ((((v13 + (("foo bar" + v15) + v13)) + "\n\t") + v15) + "\n\t"))))}
                        _ = v24
                    default:
                        if len(v2) > 0 {
                            v2[0] = v15
                        }
                    }
                    v25 := GenS3{F0: (v15 + v13)}
                    _ = v25
                }
                var v26 int64 = -1
                var v27 bool = (("hello" <= v15) && (v4 <= v4))
                _ = v26
                _ = v27
            }
            v28 := GenS4{F0: (v5 * v5), F1: 90, F2: 1}
            _ = v28
        } else {
            var v29 uint64 = 4294967295
            var v30 uint = (v5 + 0)
            fmt.Println(v1, v30, v14)
            _ = v29
            _ = v30
        }
        v6 = (v6 - (v6 & v6))
        _ = fi21
    }
    var v31 uint8 = 121
    _ = len(v2)
    if true {
        if true {
            if true {
                v32 := GenS3{F0: v13}
                switch -9223372036854775808 {
                case -9223372036854775808:
                    v33 := GenS4{F0: (v7 / ((42) | 1)), F1: v6, F2: -9223372036854775808}
                    _ = v33
                case 128:
                    v34 := GenS1{F0: true, F1: 888, F2: v31}
                    _ = v34
                }
                var v35 uint = (^(v7))
                _ = v32
                _ = v35
            }
            var v36 string = ((v15 + v15) + v15)
            _ = v36
        }
        for fi37 := 0; fi37 < 14; fi37++ {
            fmt.Println(v3, v14)
            _ = fi37
        }
        v5 += (v7 ^ v5)
    } else {
        v38 := GenS3{F0: ((((v15 + v15) + "foo bar") + ((v15 + ("\n\t" + "hello")) + v13)) + ("" + (v13 + v13)))}
        _ = v38
    }
    fmt.Println(v14, v4, v14)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
