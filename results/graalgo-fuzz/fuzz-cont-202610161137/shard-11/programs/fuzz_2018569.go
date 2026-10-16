package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 bool; F1 int16; F2 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 1.5
    var v3 int = 32768
    var v4 bool = false
    var v5 float32 = (((v2 - ((0.1 * v2) + ((float32(float64((v2 - v2))) + 1.5) - (0.1 - (float32(0.1) - v2))))) + float32(((1.5 * 1.5) + float64(((((float32(float64(0.0)) / (v2 / v2)) / float32((1.5 * float64(-2.25)))) * 0.0) + (v2 + v2)))))) + (v2 - ((0.0 * 1e10) / v2)))
    var v6 uint64 = 140
    for fi7 := 0; fi7 < 6; fi7++ {
        v3 -= 1
        delete(v1, 2)
        _ = fi7
    }
    var v8 uint64 = (^(v6))
    var v9 string = ((("a" + "a") + ("" + "")) + ("hello" + (((("a" + "\n\t") + "foo bar") + (((("\n\t" + ("a" + ("foo bar" + ((((("foo bar" + ("" + "\n\t")) + "") + ("a" + ((("foo bar" + "foo bar") + "hello") + ("foo bar" + (("" + "hello") + "\n\t"))))) + "foo bar") + ((("foo bar" + "") + "hello") + ("a" + ((((("foo bar" + "") + "foo bar") + "hello") + ("hello" + ((((("hello" + "\n\t") + "\n\t") + "hello") + "foo bar") + ("foo bar" + ("\n\t" + "hello"))))) + "\n\t"))))))) + "foo bar") + "hello") + "")) + "a")))
    v10 := GenS1{F0: uint8(v6), F1: 908}
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 uint = 515
            v14 := v1[2]
            _ = fi12
            _ = v13
            _ = v14
        }
        v2 -= (-2.25 - float32((-2.25 / ((3.14159 * 1.5) * (3.14159 / float64(v2))))))
        var v15 uint64 = v6
        _ = fi11
        _ = v15
    }
    var v16 int64 = 1
    for fi17 := 0; fi17 < 12; fi17++ {
        v5 = v2
        _ = fi17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := []int{fi18, (v3 % (((fi19 | 128)) | 1))}
            var v21 uint16 = 2
            v22 := GenS2{F0: v4, F1: 42, F2: (v9 + ("" + "foo bar"))}
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        var v23 int8 = int8(fi18)
        var v24 string = ("foo bar" + ("foo bar" + v9))
        _ = fi18
        _ = v23
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
