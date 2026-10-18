package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint16 = 65535
    var v2 float64 = ((-2.25 / 1.0) * ((3.14159 / 1.5) * (float64(float32(0.0)) * 0.0)))
    var v3 uint64 = 2
    v4 := []int{1, -266, 128}
    var v5 int = 4294967295
    v1 += 32767
    var v6 float64 = v2
    var v7 uint32 = 13
    _ = len(v4)
    if (true && (!true)) {
        if (("a" > ("" + ((("foo bar" + "\n\t") + "hello") + (("" + "foo bar") + ("hello" + ("a" + ((("hello" + "hello") + "") + "a"))))))) && true) {
            v8 := GenS1{F0: (v7 ^ (v7 >> ((255) & 63))), F1: 2, F2: v1}
            for fi9 := 0; fi9 < 5; fi9++ {
                fmt.Println(v6, fi9, v1)
                v4 = append(v4, v5)
                var v10 int8 = 0
                _ = fi9
                _ = v10
            }
            v2 = (float64(((0.1 * 1e10) / 3.14159)) - 3.14159)
            _ = v8
        }
    } else {
        v11 := GenS2{F0: false}
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            if false {
                v14 := GenS3{F0: 42}
                v1 = v1
                _ = v14
            }
            _ = fi13
        }
        var v15 float64 = 3.14159
        v4 = append(v4, 127)
        _ = fi12
        _ = v15
    }
    if (v3 >= v3) {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 int8 = -1
            var v18 uint32 = uint32(v3)
            var v19 bool = false
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
    }
    v2 += (float64((3.14159 - ((((1.5 * (0.0 * 1.5)) - (((3.14159 + (float32(1e10) * 0.0)) / (-2.25 / float32(1.5))) / (1e10 - float32(3.14159)))) * 1e10) - 3.14159))) + ((((-2.25 + (((float64(1e10) * float64(1e10)) + (v6 / 1.0)) / ((v6 - (3.14159 - 0.1)) * v2))) / 0.1) + v6) / 1.5))
    v1 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
