package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int8) int32 {
    var v5 int = 121
    switch (("hello" + "\n\t") + "") {
    case "a":
        v6 := GenS1{F0: -935, F1: (("foo bar" + "hello") + ((("" + "hello") + "") + ("a" + "foo bar")))}
        _ = v6
    case "":
        var v7 bool = ((1e10 == (1e10 + float64(-2.25))) || (!(!(false || ((543 >= 888) && false)))))
        _ = v7
    case "hello":
        v5 = v5
    }
    _ = p0
    _ = v5
    return 42
}

func main() {
    var v1 int = 1
    var v2 string = "a"
    var v3 float64 = (3.14159 * -2.25)
    var v4 int32 = 0
    var v8 int64 = 942
    v9 := []int{v1}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    v4 = v4
    _ = len(v9)
    v10 := GenS2{F0: ("" + v2)}
    for fi11 := 0; fi11 < 15; fi11++ {
        v2 = "\n\t"
        if (1 > 0) {
            var v12 int32 = v4
            _ = v12
        }
        var v13 int16 = 32767
        _ = fi11
        _ = v13
    }
    fmt.Println(v8, v8, v3)
    if (int32(v1) != (v4 % ((genFn0(13)) | 1))) {
        v14 := GenS1{F0: 1, F1: "a"}
        var v15 int8 = 20
        v16 := GenS3{F0: 241}
        _ = v14
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 4; fi17++ {
            v3 = (-2.25 * v3)
            _ = fi17
        }
    }
    v18 := GenS2{F0: v2}
    if false {
        _ = len(v9)
        if false {
            v4 -= v4
        }
        v8 += v8
    }
    v19 := GenS4{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
