package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint }
func genFn0(p0 string) int32 {
    v5 := []int{4294967295}
    p0 = p0
    p0 = "hello"
    fmt.Println(p0)
    var v6 bool = (203 != -611)
    _ = p0
    _ = v5
    _ = v6
    return 347
}
func genFn1(p0 uint16, p1 string) int8 {
    v7 := []int{-616, 42, 0}
    var v8 int32 = -773
    v8 = v8
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint = 255
    var v2 uint8 = 1
    var v3 uint32 = 2147483648
    var v4 bool = (32767 == 255)
    if true {
        v9 := GenS1{F0: 1000, F1: -128}
        _ = v9
    } else {
        var v10 int8 = 0
        var v11 uint8 = (^(v2))
        _ = v10
        _ = v11
    }
    if (("\n\t" + (("" + ("foo bar" + "\n\t")) + ("\n\t" + (("hello" + ("\n\t" + "foo bar")) + "\n\t")))) != "a") {
        v12 := GenS2{F0: 2}
        _ = v12
    } else {
        var v13 string = ("a" + "\n\t")
        var v14 int = int(v2)
        _ = v13
        _ = v14
    }
    switch -2147483648 {
    case -255:
        v3 += v3
    case -1:
        v15 := GenS2{F0: uint(v2)}
        _ = v15
    case 127:
        v16 := []int{}
        _ = v16
    }
    var v17 int32 = 42
    for fi18 := 0; fi18 < 16; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            v20 := []int{fi19, (^(fi18))}
            var v21 uint16 = 255
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    var v22 int64 = -62
    var v23 int8 = 100
    v24 := []int{814, -1}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
