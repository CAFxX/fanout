package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -218
    var v3 uint32 = 2147483647
    var v4 uint16 = 0
    var v5 string = "\n\t"
    var v6 string = v5
    fmt.Println(v5)
    var v7 uint8 = uint8(v3)
    fmt.Println(v6)
    var v8 uint = 42
    v9 := v1[(^(v3))]
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int8 = -1
        _ = fi10
        _ = v11
    }
    if false {
        var v12 string = ""
        _ = v12
    } else {
        var v13 uint16 = v4
        _ = v13
    }
    var v14 int32 = 427
    v8 -= (^(v8))
    v2 = 1
    fmt.Println(v2, v14)
    v8 -= (v8 / (((v8 + v8)) | 1))
    switch ((("hello" + v9) + v9) + "a") {
    case "":
        v7 += (^(v7))
    case "hello":
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 int32 = v14
            _ = fi15
            _ = v16
        }
    case "a":
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 uint8 = v7
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 int64 = 0
                if ((v18 | (v7 / ((0) | 1))) == (^(v18))) {
                    v7 -= 0
                    v21 := GenS1{F0: (v20 >> ((uint(v7)) & 63))}
                    _ = v21
                } else {
                    var v22 uint32 = v3
                    var v23 uint32 = v3
                    v23 = (v23 & (v22 & v3))
                    _ = v22
                    _ = v23
                }
                v24 := []int{fi17, fi19}
                _ = fi19
                _ = v20
                _ = v24
            }
            _ = fi17
            _ = v18
        }
    }
    v3 += (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
