package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
func genFn0(p0 int32, p1 uint8) int32 {
    var v4 int64 = -9223372036854775808
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := []int{fi5, (^(fi5)), (fi5 & fi5), fi5}
        _ = fi5
        _ = v6
    }
    var v7 int8 = 42
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return (p0 | 1)
}
func genFn1(p0 uint8, p1 uint32) string {
    p0 = 2
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: fi8}
        if true {
            switch "a" {
            case "\n\t":
                var v10 int = (fi8 & fi8)
                _ = v10
            case "a":
                p1 += (p1 << ((1) & 63))
            case "":
                var v11 int16 = -32768
                _ = v11
            }
            var v12 uint64 = 18446744073709551615
            for fi13 := 0; fi13 < 16; fi13++ {
                switch fi8 {
                case 0:
                    v14 := []int{fi13, (fi8 - fi8), (^(fi8)), 42}
                    var v15 uint16 = 32767
                    p1 = (p1 ^ p1)
                    _ = v14
                    _ = v15
                case -576:
                    fmt.Println(v12, p0)
                case 1:
                    switch (fi13 ^ fi13) {
                    case 953:
                        var v16 int32 = 2147483647
                        _ = v16
                    case 127:
                        var v17 int16 = 1000
                        _ = v17
                    default:
                        v18 := []int{(fi8 - (fi8 + (fi13 ^ 0))), fi13, -286}
                        _ = v18
                    }
                }
                var v19 string = ""
                var v20 bool = false
                _ = fi13
                _ = v19
                _ = v20
            }
            _ = v12
        } else {
            var v21 uint16 = 0
            if true {
                switch fi8 {
                case 2147483647:
                    v22 := []int{fi8, fi8, 9223372036854775807, (fi8 ^ -2147483648)}
                    _ = v22
                }
            } else {
                v23 := GenS1{F0: (fi8 + (fi8 * (-(fi8))))}
                _ = v23
            }
            v24 := GenS1{F0: fi8}
            _ = v21
            _ = v24
        }
        switch "hello" {
        case "foo bar":
            fmt.Println(fi8, fi8)
        case "\n\t":
            v25 := []int{fi8, (fi8 / (((fi8 + (fi8 / (((fi8 & fi8)) | 1)))) | 1))}
            _ = v25
        }
        _ = fi8
        _ = v9
    }
    v26 := []int{32768, 4294967295, 9223372036854775807, 128}
    if ((!(!false)) && true) {
        p0 -= (^(p0))
        v27 := GenS1{F0: -35}
        _ = v27
    } else {
        if (uint64(p0) != 4294967295) {
            var v28 int64 = -9223372036854775808
            _ = v28
        } else {
            _ = len(v26)
        }
        var v29 uint8 = uint8(p1)
        _ = v29
    }
    var v30 float32 = -2.25
    _ = p0
    _ = p1
    _ = v26
    _ = v30
    return "hello"
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{("foo bar" + "\n\t"), (("a" + "") + "")}
    var v3 string = "\n\t"
    for fi31 := 0; fi31 < 12; fi31++ {
        v32 := GenS1{F0: (fi31 << ((18446744073709551615) & 63))}
        var v33 int32 = -2147483648
        _ = fi31
        _ = v32
        _ = v33
    }
    v3 = "foo bar"
    v34 := GenS1{F0: -1}
    for fi35 := 0; fi35 < 3; fi35++ {
        var v36 uint8 = 1
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        v3 = v3
        var v38 float32 = 1.5
        v1[v3] = (^(fi37))
        _ = fi37
        _ = v38
    }
    v39 := GenS2{F0: 1}
    v3 = (v3 + v3)
    fmt.Println(v3)
    var v40 string = (v3 + v3)
    v40 = v3
    delete(v1, "foo bar")
    v41 := GenS1{F0: 857}
    v2 = append(v2, "foo bar")
    var v42 bool = (100 == -128)
    var v43 uint64 = 42
    v44 := GenS2{F0: 553}
    if false {
        fmt.Println(v42)
        var v45 int64 = 0
        _ = v45
    } else {
        v40 = (((v40 + v3) + ((((v3 + "a") + v40) + ("\n\t" + v3)) + v3)) + v3)
    }
    v3 = ("hello" + (v3 + (v3 + v3)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
