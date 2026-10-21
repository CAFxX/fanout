package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int32, p1 uint32, p2 uint16) uint16 {
    var v5 uint8 = 215
    var v6 uint16 = 0
    p1 -= (p1 << ((4294967295) & 63))
    p2 -= v6
    p2 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 353
}
func genFn1(p0 string) bool {
    if false {
        if false {
            switch (p0 + "a") {
            case "":
                for fi7 := 0; fi7 < 19; fi7++ {
                    if false {
                        v8 := []int{fi7, fi7, (-(fi7))}
                        p0 = ("a" + p0)
                        _ = v8
                    }
                    if false {
                        switch p0 {
                        case "foo bar":
                            for fi9 := 0; fi9 < 11; fi9++ {
                                var v10 uint32 = 1
                                _ = fi9
                                _ = v10
                            }
                        case "a":
                            v11 := []int{fi7}
                            _ = v11
                        }
                        var v12 string = "foo bar"
                        _ = v12
                    }
                    v13 := GenS1{F0: (fi7 >> ((4294967295) & 63))}
                    _ = fi7
                    _ = v13
                }
            case "foo bar":
                switch -1 {
                case 42:
                    var v14 uint = 42
                    _ = v14
                case 9223372036854775807:
                    var v15 bool = false
                    _ = v15
                }
            case "a":
                var v16 uint16 = genFn0(-662, 0, 0)
                _ = v16
            default:
                var v17 uint64 = 1
                _ = v17
            }
            switch 0 {
            case 339:
                p0 = p0
            case 127:
                var v18 uint64 = 4294967295
                _ = v18
            case 971:
                v19 := []int{4294967295, 32768, 4294967295}
                _ = v19
            default:
                v20 := []int{922, 42, -745, -210}
                _ = v20
            }
        }
    }
    fmt.Println(p0)
    _ = p0
    return (136 != 192)
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    v4 := []string{(("\n\t" + ("hello" + "\n\t")) + "\n\t"), "foo bar"}
    v21 := GenS1{F0: 102}
    var v22 uint32 = 2
    v23 := GenS1{F0: -9223372036854775808}
    v22 = (v22 + 0)
    var v24 uint = 762
    fmt.Println(v24, v22, v1)
    if ("a" >= ("foo bar" + "foo bar")) {
        switch 4294967295 {
        case -2147483648:
            fmt.Println(v1)
        }
        var v25 int = 32768
        _ = v25
    }
    var v26 int8 = 107
    if (((!(v22 <= 2)) && (!false)) && false) {
        delete(v3, 2147483648)
        var v27 float32 = v1
        for fi28 := 0; fi28 < 5; fi28++ {
            v22 = v22
            v27 = 3.14159
            _ = fi28
        }
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
}
