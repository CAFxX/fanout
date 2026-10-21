package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint64, p2 uint64) int16 {
    var v7 bool = true
    v8 := GenS1{F0: 42, F1: "a"}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -32768
}
func genFn1() int64 {
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: 0, F1: ("foo bar" + ("" + "a"))}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS1{F0: 243, F1: "foo bar"}
        v13 := []int{fi11, fi11}
        _ = fi11
        _ = v12
        _ = v13
    }
    return -580
}

func main() {
    var v1 uint8 = 127
    var v2 float64 = 1.5
    var v3 int16 = -441
    v4 := make(map[uint32]bool)
    v5 := []string{(("hello" + "\n\t") + (("hello" + ("\n\t" + ("" + "hello"))) + (("a" + "\n\t") + ((("foo bar" + "\n\t") + "a") + ("\n\t" + ((("" + "\n\t") + "\n\t") + ((((("\n\t" + "a") + ("hello" + "foo bar")) + "\n\t") + ("foo bar" + "hello")) + "foo bar"))))))), (("" + "hello") + ((("a" + "foo bar") + "a") + "hello"))}
    var v6 string = ""
    var v14 uint8 = 127
    var v15 uint16 = uint16(v1)
    v4[1] = true
    if false {
        v14 += (v14 - (v14 & v1))
        if len(v5) > 0 {
            v5[0] = (((v6 + "hello") + "foo bar") + (v6 + v6))
        }
    } else {
        var v16 int16 = v3
        _ = v16
    }
    v17 := GenS2{F0: ""}
    var v18 uint = uint(v1)
    v19 := GenS3{F0: (^(v18))}
    var v20 int32 = -2147483648
    if true {
        v21 := GenS4{F0: v14}
        v22 := GenS1{F0: 127, F1: (v6 + ("hello" + (v6 + "hello")))}
        _ = v21
        _ = v22
    }
    v15 += v15
    var v23 float64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
