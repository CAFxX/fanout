package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 string; F2 uint8 }
type GenS3 struct { F0 int64; F1 int32; F2 int16 }
func genFn0(p0 uint) int8 {
    var v5 bool = (143 <= 0)
    if v5 {
        switch 673 {
        case 9223372036854775807:
            var v6 int64 = 1
            _ = v6
        default:
            p0 = (p0 / (((p0 & 196)) | 1))
        }
        v7 := GenS1{F0: -1}
        _ = v7
    } else {
        if (4294967295 <= 2147483647) {
            p0 += p0
            for fi8 := 0; fi8 < 10; fi8++ {
                switch "a" {
                case "foo bar":
                    v9 := []int{fi8, fi8, fi8}
                    _ = v9
                case "hello":
                    p0 = p0
                default:
                    v10 := GenS2{F0: 2, F1: ("" + "hello"), F2: 42}
                    _ = v10
                }
                p0 = p0
                _ = fi8
            }
        } else {
            var v11 int32 = -458
            _ = v11
        }
    }
    v5 = (!true)
    if false {
        v12 := GenS1{F0: 0}
        _ = v12
    }
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 bool = false
    var v2 uint8 = 128
    v3 := []string{"\n\t", ("foo bar" + ("\n\t" + "foo bar")), ((("\n\t" + ("\n\t" + (("a" + ("hello" + ("" + ""))) + "a"))) + (("hello" + "\n\t") + "hello")) + "\n\t")}
    var v4 int16 = -690
    v4 = v4
    var v13 uint = 18446744073709551615
    if (1e10 <= 1.5) {
        var v14 int64 = -9223372036854775808
        _ = v14
    } else {
        var v15 float32 = -2.25
        v1 = v1
        for fi16 := 0; fi16 < 5; fi16++ {
            if (((4294967295 != 1) || (v1 && (!(!(-288 != -1))))) || false) {
                fmt.Println(v1)
                fmt.Println(v15, v4)
                fmt.Println(v1, v1)
            } else {
                if (v1 || v1) {
                    v17 := GenS3{F0: 490, F1: -1, F2: -935}
                    var v18 bool = false
                    _ = v17
                    _ = v18
                }
                if (!v1) {
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                    var v19 uint16 = 65535
                    _ = v19
                }
            }
            _ = fi16
        }
        _ = v15
    }
    _ = len(v3)
    v1 = true
    v1 = v1
    if (0 > 891) {
        v13 += 529
    } else {
        v2 = v2
    }
    v20 := GenS3{F0: -474, F1: 42, F2: 42}
    fmt.Println(v13)
    var v21 uint16 = 86
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
