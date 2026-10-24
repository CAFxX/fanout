package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
func genFn0(p0 uint8, p1 int8, p2 uint) int8 {
    v6 := GenS1{F0: 144, F1: p1, F2: 166}
    var v7 int = 32768
    var v8 float64 = -2.25
    v9 := GenS2{F0: "foo bar", F1: 403}
    v7 -= v7
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (p1 * (p1 ^ -1))
}
func genFn1(p0 string) uint64 {
    var v10 int64 = 587
    var v11 int8 = 127
    v12 := []int{257, int(v10), 127, 1}
    v10 += v10
    v11 = (v11 ^ -128)
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 string = "a"
    var v2 int64 = 1
    v3 := []string{(v1 + ("\n\t" + (v1 + ((("" + (v1 + (v1 + "a"))) + ((v1 + ("" + (v1 + ("hello" + ((v1 + v1) + "hello"))))) + ((v1 + (("a" + v1) + "")) + "foo bar"))) + ((v1 + ((v1 + "\n\t") + v1)) + (v1 + "a")))))), v1}
    var v4 int32 = -179
    var v5 int = -1
    v1 = v1
    v13 := GenS3{F0: v2}
    v14 := GenS4{F0: -9223372036854775808}
    var v15 int64 = 1
    var v16 int8 = 127
    v16 -= 111
    v1 = v1
    switch v1 {
    case "hello":
        for fi17 := 0; fi17 < 7; fi17++ {
            if (0.0 < float64(0.1)) {
                for fi18 := 0; fi18 < 10; fi18++ {
                    v4 += v4
                    v15 = (v15 >> ((196) & 63))
                    v1 = ("\n\t" + v1)
                    _ = fi18
                }
            }
            v3 = append(v3, "foo bar")
            _ = fi17
        }
    }
    var v19 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
