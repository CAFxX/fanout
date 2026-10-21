package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int, p1 uint8, p2 uint16) uint8 {
    var v4 int64 = 206
    for fi5 := 0; fi5 < 10; fi5++ {
        switch "foo bar" {
        case "\n\t":
            if (!true) {
                var v6 int8 = 65
                _ = v6
            } else {
                v7 := GenS1{F0: 860}
                switch -503 {
                case -394:
                    v8 := []int{fi5, (^(p0)), fi5}
                    _ = v8
                }
                _ = v7
            }
        default:
            var v9 float32 = (-2.25 + 1.5)
            _ = v9
        }
        _ = fi5
    }
    v10 := GenS2{F0: p2}
    switch ("" + ((("foo bar" + ("\n\t" + ("a" + ("hello" + "\n\t")))) + ("foo bar" + ("a" + ((("" + ((((("foo bar" + ("hello" + "foo bar")) + "a") + ("" + "\n\t")) + "a") + "foo bar")) + "\n\t") + ("" + "\n\t"))))) + "foo bar")) {
    case "\n\t":
        if true {
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 int32 = 42
                _ = fi11
                _ = v12
            }
            var v13 bool = (-2.25 == (1.5 / float32(float64(float32(0.1)))))
            _ = v13
        }
    default:
        v14 := GenS3{F0: 155}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v4
    return (p1 >> ((234) & 63))
}
func genFn1() int32 {
    v15 := []int{}
    var v16 float32 = -2.25
    if len(v15) > 0 {
        v15[0] = 1
    }
    v16 = (0.1 / (0.0 / v16))
    _ = v15
    _ = v16
    return 0
}

func main() {
    var v1 uint8 = 255
    var v2 uint32 = 0
    var v3 float64 = 3.14159
    if true {
        if true {
            v2 = (^(v2))
        }
    }
    var v17 int32 = -2147483648
    var v18 uint32 = (v2 >> ((0) & 63))
    v19 := GenS2{F0: 42}
    switch "foo bar" {
    case "\n\t":
        if true {
            v20 := []int{128}
            v20 = append(v20, -503)
            v21 := GenS3{F0: 18446744073709551615}
            _ = v20
            _ = v21
        }
    case "hello":
        v22 := GenS1{F0: 329}
        _ = v22
    }
    var v23 int32 = genFn1()
    v24 := GenS3{F0: 42}
    var v25 uint64 = 226
    var v26 bool = true
    fmt.Println(v25)
    v27 := GenS3{F0: v25}
    v28 := []int{4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
