package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int; F2 bool }
type GenS4 struct { F0 uint64 }
func genFn0(p0 bool) uint32 {
    p0 = (42 <= 32767)
    v5 := []int{485}
    switch ("" + ("a" + "a")) {
    case "":
        for fi6 := 0; fi6 < 5; fi6++ {
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 int64 = 0
                var v9 uint16 = 1
                switch fi6 {
                case 128:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
    }
    _ = p0
    _ = v5
    return 65535
}

func main() {
    var v1 uint8 = 0
    var v2 string = (("hello" + "") + ("\n\t" + "hello"))
    var v3 bool = ((v2 + v2) > ("\n\t" + ("" + (("a" + "") + v2))))
    var v4 uint8 = 127
    fmt.Println(v1, v4)
    v10 := GenS1{F0: -2147483648, F1: 0}
    var v11 uint = 0
    fmt.Println(v11, v3, v2)
    for fi12 := 0; fi12 < 14; fi12++ {
        if (!v3) {
            for fi13 := 0; fi13 < 13; fi13++ {
                if v3 {
                    fmt.Println(fi12, v1)
                    switch v2 {
                    case "hello":
                        var v14 string = v2
                        _ = v14
                    case "\n\t":
                        v1 += (v1 % ((v1) | 1))
                    case "":
                        var v15 int64 = -682
                        _ = v15
                    default:
                        v2 = (v2 + v2)
                    }
                } else {
                    if (v1 <= (v4 | (v4 >> ((1) & 63)))) {
                        v11 = (v11 | v11)
                        v4 += (v4 - 127)
                        var v16 string = ((v2 + "\n\t") + "foo bar")
                        _ = v16
                    }
                }
                v2 = v2
                v17 := GenS2{F0: 0, F1: 1}
                _ = fi13
                _ = v17
            }
            v18 := GenS3{F0: 795, F1: fi12, F2: v3}
            for fi19 := 0; fi19 < 19; fi19++ {
                var v20 uint8 = (v1 + (v1 % ((74) | 1)))
                var v21 int64 = 0
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = v18
        } else {
            v22 := []int{fi12, (fi12 + 128)}
            _ = v22
        }
        _ = fi12
    }
    switch 625 {
    case 128:
        for fi23 := 0; fi23 < 8; fi23++ {
            v24 := []int{fi23}
            v25 := GenS2{F0: (^(v11)), F1: 42}
            v26 := GenS4{F0: 683}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
    }
    var v27 string = ((v2 + "") + v2)
    v28 := GenS2{F0: (v11 - v11), F1: 210}
    var v29 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
