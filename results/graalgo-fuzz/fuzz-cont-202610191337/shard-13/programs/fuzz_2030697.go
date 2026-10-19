package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint8, p1 uint32) uint {
    v6 := GenS1{F0: ""}
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 float32 = 1.5
    var v2 uint = 255
    v3 := make(map[string]string)
    v4 := []int{32768, 128, 1}
    var v5 float32 = ((float32((3.14159 + 0.1)) - ((v1 * v1) * v1)) + (((float32(-2.25) - -2.25) + float32((1e10 * 1.5))) - (((v1 + (v1 - float32((0.0 - -2.25)))) + 1e10) * ((v1 / 1.0) - ((float32(-2.25) - float32((0.1 / 1.0))) * v1)))))
    delete(v3, ("hello" + "foo bar"))
    v7 := GenS2{F0: -17}
    var v8 uint32 = 65535
    v5 = 3.14159
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: ("" + ("hello" + "foo bar"))}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := GenS2{F0: -6}
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: ("foo bar" + "hello")}
    v14 := GenS3{F0: 42, F1: 122}
    delete(v3, (("hello" + "") + ("hello" + "a")))
    v5 -= 3.14159
    v1 = (v1 * v5)
    if false {
        if (!((v8 == v8) && true)) {
            var v15 uint32 = uint32(v2)
            _ = v15
        }
        v8 += 2
    } else {
        var v16 uint32 = v8
        if false {
            var v17 string = "\n\t"
            fmt.Println(v5, v2)
            _ = v17
        } else {
            var v18 int32 = 42
            v19 := GenS2{F0: 10}
            _ = v18
            _ = v19
        }
        _ = v16
    }
    var v20 uint16 = 510
    var v21 string = (("foo bar" + ("\n\t" + "foo bar")) + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
