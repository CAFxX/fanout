package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint16 }

func main() {
    var v1 string = ("\n\t" + "")
    v2 := []int{580, -1, 799}
    var v3 int8 = 42
    var v4 string = v1
    if false {
        switch 1 {
        case -2147483648:
            var v5 int32 = -760
            _ = v5
        case 937:
            if ((388 < 18446744073709551615) || false) {
                fmt.Println(v4)
                v6 := GenS1{F0: 0, F1: 2147483647, F2: 1}
                for fi7 := 0; fi7 < 9; fi7++ {
                    var v8 uint16 = 736
                    v3 -= 42
                    var v9 uint32 = 824
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                _ = v6
            } else {
                var v10 int64 = -1
                v11 := GenS1{F0: 2, F1: 268, F2: 0}
                _ = len(v2)
                _ = v10
                _ = v11
            }
        default:
            var v12 float32 = float32(1.5)
            _ = v12
        }
        switch ("\n\t" + "a") {
        case "":
            for fi13 := 0; fi13 < 2; fi13++ {
                v3 = -1
                _ = fi13
            }
        case "a":
            v1 = v1
        case "\n\t":
            for fi14 := 0; fi14 < 13; fi14++ {
                if true {
                    v15 := GenS1{F0: 502, F1: (^(fi14)), F2: 1}
                    _ = v15
                } else {
                    var v16 uint32 = 2147483647
                    switch (fi14 % (((fi14 + (-(fi14)))) | 1)) {
                    case 2147483647:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -968:
                        if (true || false) {
                            var v17 string = v4
                            var v18 int64 = 42
                            var v19 float64 = 3.14159
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                    }
                    _ = v16
                }
                var v20 uint16 = 42
                v20 -= v20
                _ = fi14
                _ = v20
            }
        }
    }
    var v21 bool = ((v3 % (((v3 % (((^(v3))) | 1))) | 1)) <= 42)
    if ((1 > 0) && v21) {
        var v22 float64 = -2.25
        _ = v22
    }
    v23 := GenS1{F0: 65535, F1: -1, F2: 32767}
    var v24 int32 = -5
    for fi25 := 0; fi25 < 16; fi25++ {
        if v21 {
            var v26 int = -1
            fmt.Println(fi25, v4, fi25)
            fmt.Println(v26, v4, v21)
            _ = v26
        }
        _ = fi25
    }
    v2 = append(v2, 32767)
    var v27 int = 32768
    var v28 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
